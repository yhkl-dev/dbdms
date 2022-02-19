package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetDatabaseByID(id int) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, name, host, port, username, schema, comment, created_at, updated_at from t_databases where id = $1`
	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var database Database
	rows.Scan(
		&database.ID,
		&database.Name,
		&database.Host,
		&database.Port,
		&database.Username,
		&database.Schema,
		&database.Comment,
		&database.CreatedAt,
		&database.UpdatedAt,
	)
	if database.ID == 0 {
		return nil, nil
	}
	return &database, nil
}

func (m *DBModel) ListDatabases(genre ...int) ([]*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, name, host, port, username, schema, comment, created_at, updated_at from t_databases`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []*Database
	for rows.Next() {
		var database Database
		err := rows.Scan(
			&database.ID,
			&database.Name,
			&database.Host,
			&database.Port,
			&database.Username,
			&database.Schema,
			&database.Comment,
			&database.CreatedAt,
			&database.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		databases = append(databases, &database)
	}
	return databases, nil
}

func (m *DBModel) CreateDatabase(data Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into t_databases 
		(name, host, port, username, password, schema, comment, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		data.Name, data.Host, data.Port, data.Username, data.Password,
		data.Schema, data.Comment, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) UpdateDatabase(data Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `
		update t_databases set 
			name = $1,
			host = $2,
			port = $3,
			username = $4,
			password = $5,
			schema = $6,
			comment = $7,
			created_at = $8,
			updated_at = $9
		where 
			id = $10
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		data.Name, data.Host, data.Port, data.Username, data.Password,
		data.Schema, data.Comment, data.CreatedAt, data.UpdatedAt, data.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) DeleteDatabase(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `
		delete from t_databases 
		where 
			id = $1
	`
	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
