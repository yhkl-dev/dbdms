package mysql

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"dbdms/tbls/drivers"
	"dbdms/tbls/schema"
	"github.com/pkg/errors"
)

var reFK = regexp.MustCompile(`FOREIGN KEY \((.+)\) REFERENCES ([^\s]+)\s?\((.+)\)`)
var reAI = regexp.MustCompile(`AUTO_INCREMENT=[\d]+`)

// Mysql struct
type Mysql struct {
	db                *sql.DB
	showAutoIncrement bool
}

func ShowAutoIcrrement() drivers.Option {
	return func(d drivers.Driver) error {
		switch d := d.(type) {
		case *Mysql:
			d.showAutoIncrement = true
		}
		return nil
	}
}

// New return new Mysql
func New(db *sql.DB, opts ...drivers.Option) (*Mysql, error) {
	m := &Mysql{
		db: db,
	}
	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Analyze MySQL database schema
func (m *Mysql) Analyze(s *schema.Schema) error {
	d, err := m.Info()
	if err != nil {
		return errors.WithStack(err)
	}
	s.Driver = d

	// tables and comments
	tableRows, err := m.db.Query(`
SELECT table_name, table_type, table_comment FROM information_schema.tables WHERE table_schema = ?;`, s.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	defer tableRows.Close()

	relations := []*schema.Relation{}

	tables := []*schema.Table{}
	for tableRows.Next() {
		var (
			tableName    string
			tableType    string
			tableComment string
		)
		err := tableRows.Scan(&tableName, &tableType, &tableComment)
		if err != nil {
			return errors.WithStack(err)
		}
		table := &schema.Table{
			Name:    tableName,
			Type:    tableType,
			Comment: tableComment,
		}

		// table definition
		if tableType == "BASE TABLE" {
			tableDefRows, err := m.db.Query(fmt.Sprintf("SHOW CREATE TABLE `%s`", tableName))
			if err != nil {
				return errors.WithStack(err)
			}
			defer tableDefRows.Close()
			for tableDefRows.Next() {
				var (
					tableName string
					tableDef  string
				)
				err := tableDefRows.Scan(&tableName, &tableDef)
				if err != nil {
					return errors.WithStack(err)
				}
				if m.showAutoIncrement {
					table.Def = tableDef
				} else {
					table.Def = reAI.ReplaceAllLiteralString(tableDef, "AUTO_INCREMENT=[Redacted by tbls]")
				}
			}
		}

		// view definition
		if tableType == "VIEW" {
			viewDefRows, err := m.db.Query(`
SELECT view_definition FROM information_schema.views
WHERE table_schema = ?
AND table_name = ?;
		`, s.Name, tableName)
			if err != nil {
				return errors.WithStack(err)
			}
			defer viewDefRows.Close()
			for viewDefRows.Next() {
				var tableDef string
				err := viewDefRows.Scan(&tableDef)
				if err != nil {
					return errors.WithStack(err)
				}
				table.Def = fmt.Sprintf("CREATE VIEW %s AS (%s)", tableName, tableDef)
			}
		}

		// indexes
		indexRows, err := m.db.Query(`
SELECT
(CASE WHEN s.index_name='PRIMARY' AND s.non_unique=0 THEN 'PRIMARY KEY'
      WHEN s.index_name!='PRIMARY' AND s.non_unique=0 THEN 'UNIQUE KEY'
      WHEN s.non_unique=1 THEN 'KEY'
      ELSE null
  END) AS key_type,
s.index_name, GROUP_CONCAT(s.column_name ORDER BY s.seq_in_index SEPARATOR ', '), s.index_type
FROM information_schema.statistics AS s
LEFT JOIN information_schema.columns AS c ON s.table_schema = c.table_schema AND s.table_name = c.table_name AND s.column_name = c.column_name
WHERE s.table_name = c.table_name
AND s.table_schema = ?
AND s.table_name = ?
GROUP BY key_type, s.table_name, s.index_name, s.index_type`, s.Name, tableName)
		if err != nil {
			return errors.WithStack(err)
		}
		defer indexRows.Close()

		indexes := []*schema.Index{}
		for indexRows.Next() {
			var (
				indexKeyType    string
				indexName       string
				indexColumnName string
				indexType       string
				indexDef        string
			)
			err = indexRows.Scan(&indexKeyType, &indexName, &indexColumnName, &indexType)
			if err != nil {
				return errors.WithStack(err)
			}

			if indexKeyType == "PRIMARY KEY" {
				indexDef = fmt.Sprintf("%s (%s) USING %s", indexKeyType, indexColumnName, indexType)
			} else {
				indexDef = fmt.Sprintf("%s %s (%s) USING %s", indexKeyType, indexName, indexColumnName, indexType)
			}

			index := &schema.Index{
				Name:    indexName,
				Def:     indexDef,
				Table:   &table.Name,
				Columns: strings.Split(indexColumnName, ", "),
			}
			indexes = append(indexes, index)
		}
		table.Indexes = indexes

		// constraints
		constraintRows, err := m.db.Query(`
SELECT
  kcu.constraint_name,
  sub.costraint_type,
  GROUP_CONCAT(kcu.column_name ORDER BY kcu.ordinal_position, position_in_unique_constraint SEPARATOR ', ') AS column_name,
  kcu.referenced_table_name,
  GROUP_CONCAT(kcu.referenced_column_name ORDER BY kcu.ordinal_position, position_in_unique_constraint SEPARATOR ', ') AS referenced_column_name
FROM information_schema.key_column_usage AS kcu
LEFT JOIN information_schema.columns AS c ON kcu.table_schema = c.table_schema AND kcu.table_name = c.table_name AND kcu.column_name = c.column_name
LEFT JOIN
  (
   SELECT
   kcu.table_schema,
   kcu.table_name,
   kcu.constraint_name,
   kcu.column_name,
   (CASE WHEN kcu.referenced_table_name IS NOT NULL THEN 'FOREIGN KEY'
        WHEN c.column_key = 'PRI' AND kcu.constraint_name = 'PRIMARY' THEN 'PRIMARY KEY'
        WHEN c.column_key = 'PRI' AND kcu.constraint_name != 'PRIMARY' THEN 'UNIQUE'
        WHEN c.column_key = 'UNI' THEN 'UNIQUE'
        WHEN c.column_key = 'MUL' THEN 'UNIQUE'
        ELSE 'UNKNOWN'
   END) AS costraint_type
   FROM information_schema.key_column_usage AS kcu
   LEFT JOIN information_schema.columns AS c ON kcu.table_schema = c.table_schema AND kcu.table_name = c.table_name AND kcu.column_name = c.column_name
   WHERE kcu.table_name = ?
   AND kcu.ordinal_position = 1
  ) AS sub
ON kcu.constraint_name = sub.constraint_name AND kcu.table_schema = sub.table_schema AND kcu.table_name = sub.table_name
WHERE kcu.table_schema= ?
   AND kcu.table_name = ?
GROUP BY kcu.constraint_name, sub.costraint_type, kcu.referenced_table_name`, tableName, s.Name, tableName)
		if err != nil {
			return errors.WithStack(err)
		}
		defer constraintRows.Close()

		constraints := []*schema.Constraint{}
		for constraintRows.Next() {
			var (
				constraintName          string
				constraintType          string
				constraintColumnName    string
				constraintRefTableName  sql.NullString
				constraintRefColumnName sql.NullString
				constraintDef           string
			)
			err = constraintRows.Scan(&constraintName, &constraintType, &constraintColumnName, &constraintRefTableName, &constraintRefColumnName)
			if err != nil {
				return errors.WithStack(err)
			}
			switch constraintType {
			case "PRIMARY KEY":
				constraintDef = fmt.Sprintf("PRIMARY KEY (%s)", constraintColumnName)
			case "UNIQUE":
				constraintDef = fmt.Sprintf("UNIQUE KEY %s (%s)", constraintName, constraintColumnName)
			case "FOREIGN KEY":
				constraintType = schema.TypeFK
				constraintDef = fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s (%s)", constraintColumnName, constraintRefTableName.String, constraintRefColumnName.String)
				relation := &schema.Relation{
					Table: table,
					Def:   constraintDef,
				}
				relations = append(relations, relation)
			case "UNKNOWN":
				constraintDef = fmt.Sprintf("UNKNOWN CONSTRAINT (%s) (%s) (%s)", constraintColumnName, constraintRefTableName.String, constraintRefColumnName.String)
			}

			constraint := &schema.Constraint{
				Name:    constraintName,
				Type:    constraintType,
				Def:     constraintDef,
				Table:   &table.Name,
				Columns: strings.Split(constraintColumnName, ", "),
			}
			if constraintRefTableName.String != "" {
				constraint.ReferenceTable = &constraintRefTableName.String
				constraint.ReferenceColumns = strings.Split(constraintRefColumnName.String, ", ")
			}

			constraints = append(constraints, constraint)
		}
		table.Constraints = constraints

		// triggers
		triggerRows, err := m.db.Query(`
SELECT
  trigger_name,
  action_timing,
  event_manipulation,
  event_object_table,
  action_orientation,
  action_statement
FROM information_schema.triggers
WHERE event_object_schema = ?
AND event_object_table = ?
`, s.Name, tableName)
		if err != nil {
			return errors.WithStack(err)
		}
		defer triggerRows.Close()
		triggers := []*schema.Trigger{}
		for triggerRows.Next() {
			var (
				triggerName              string
				triggerActionTiming      string
				triggerEventManipulation string
				triggerEventObjectTable  string
				triggerActionOrientation string
				triggerActionStatement   string
				triggerDef               string
			)
			err = triggerRows.Scan(&triggerName, &triggerActionTiming, &triggerEventManipulation, &triggerEventObjectTable, &triggerActionOrientation, &triggerActionStatement)
			if err != nil {
				return errors.WithStack(err)
			}
			triggerDef = fmt.Sprintf("CREATE TRIGGER %s %s %s ON %s\nFOR EACH %s\n%s", triggerName, triggerActionTiming, triggerEventManipulation, triggerEventObjectTable, triggerActionOrientation, triggerActionStatement)
			trigger := &schema.Trigger{
				Name: triggerName,
				Def:  triggerDef,
			}
			triggers = append(triggers, trigger)
		}
		table.Triggers = triggers

		// columns and comments
		columnRows, err := m.db.Query(`
SELECT column_name, column_default, is_nullable, column_type, column_comment
FROM information_schema.columns
WHERE table_schema = ? AND table_name = ? ORDER BY ordinal_position`, s.Name, tableName)
		if err != nil {
			return errors.WithStack(err)
		}
		defer columnRows.Close()
		columns := []*schema.Column{}
		for columnRows.Next() {
			var (
				columnName    string
				columnDefault sql.NullString
				isNullable    string
				columnType    string
				columnComment sql.NullString
			)
			err = columnRows.Scan(&columnName, &columnDefault, &isNullable, &columnType, &columnComment)
			if err != nil {
				return errors.WithStack(err)
			}
			column := &schema.Column{
				Name:     columnName,
				Type:     columnType,
				Nullable: convertColumnNullable(isNullable),
				Default:  columnDefault,
				Comment:  columnComment.String,
			}

			columns = append(columns, column)
		}
		table.Columns = columns

		tables = append(tables, table)
	}

	s.Tables = tables

	// Relations
	for _, r := range relations {
		result := reFK.FindAllStringSubmatch(r.Def, -1)
		strColumns := strings.Split(result[0][1], ", ")
		strParentTable := result[0][2]
		strParentColumns := strings.Split(result[0][3], ", ")
		for _, c := range strColumns {
			column, err := r.Table.FindColumnByName(c)
			if err != nil {
				return err
			}
			r.Columns = append(r.Columns, column)
			column.ParentRelations = append(column.ParentRelations, r)
		}
		parentTable, err := s.FindTableByName(strParentTable)
		if err != nil {
			return err
		}
		r.ParentTable = parentTable
		for _, c := range strParentColumns {
			column, err := parentTable.FindColumnByName(c)
			if err != nil {
				return err
			}
			r.ParentColumns = append(r.ParentColumns, column)
			column.ChildRelations = append(column.ChildRelations, r)
		}
	}

	s.Relations = relations

	return nil
}

// Info return schema.Driver
func (m *Mysql) Info() (*schema.Driver, error) {
	var v string
	row := m.db.QueryRow(`SELECT version();`)
	err := row.Scan(&v)
	if err != nil {
		return nil, err
	}

	d := &schema.Driver{
		Name:            "mysql",
		DatabaseVersion: v,
	}
	return d, nil
}

func convertColumnNullable(str string) bool {
	if str == "NO" {
		return false
	}
	return true
}
