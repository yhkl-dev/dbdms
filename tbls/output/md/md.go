package md

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"dbdms/tbls/config"
	"dbdms/tbls/output"
	"dbdms/tbls/schema"
	"github.com/gobuffalo/packr/v2"
	"github.com/mattn/go-runewidth"
	"github.com/pkg/errors"
	"github.com/pmezard/go-difflib/difflib"
)

// Md struct
type Md struct {
	config *config.Config
	er     bool
	box    *packr.Box
}

// New return Md
func New(c *config.Config, er bool) *Md {
	return &Md{
		config: c,
		er:     er,
		box:    packr.New("md", "./templates"),
	}
}

// OutputSchema output .md format for all tables.
func (m *Md) OutputSchema(wr io.Writer, s *schema.Schema) (data map[string]interface{}, err error) {
	ts, err := m.box.FindString("index.md.tmpl")
	if err != nil {
		return data, errors.WithStack(err)
	}
	tmpl := template.Must(template.New("index").Funcs(output.Funcs(&m.config.MergedDict)).Parse(ts))
	templateData := m.makeSchemaTemplateData(s, m.config.Format.Adjust)
	templateData["er"] = m.er
	templateData["erFormat"] = m.config.ER.Format
	err = tmpl.Execute(wr, templateData)
	if err != nil {
		return data, errors.WithStack(err)
	}
	return templateData, nil
}

// OutputTable output md format for table.
func (m *Md) OutputTable(wr io.Writer, t *schema.Table) error {
	ts, err := m.box.FindString("table.md.tmpl")
	if err != nil {
		return errors.WithStack(err)
	}
	tmpl := template.Must(template.New(t.Name).Funcs(output.Funcs(&m.config.MergedDict)).Parse(ts))
	templateData := m.makeTableTemplateData(t, m.config.Format.Adjust)
	templateData["er"] = m.er
	templateData["erFormat"] = m.config.ER.Format
	fmt.Println(">>>>>>>>>>>>", m.config.ER.Format)

	err = tmpl.Execute(wr, templateData)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Output generate markdown files.
func Output(s *schema.Schema, c *config.Config, force bool) (e error) {
	docPath := c.DocPath

	fullPath, err := filepath.Abs(docPath)
	if err != nil {
		return errors.WithStack(err)
	}

	if !force && outputExists(s, fullPath) {
		return errors.New("output files already exists")
	}

	err = os.MkdirAll(fullPath, 0755) // #nosec
	if err != nil {
		return errors.WithStack(err)
	}

	// README.md
	file, err := os.Create(filepath.Join(fullPath, "README.md"))
	defer func() {
		err := file.Close()
		if err != nil {
			e = err
		}
	}()
	if err != nil {
		return errors.WithStack(err)
	}
	er := false
	if _, err := os.Lstat(filepath.Join(fullPath, fmt.Sprintf("schema.%s", c.ER.Format))); err == nil {
		er = true
	}

	md := New(c, er)

	_, err = md.OutputSchema(file, s)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("%s\n", filepath.Join(docPath, "README.md"))

	// tables
	for _, t := range s.Tables {
		file, err := os.Create(filepath.Join(fullPath, fmt.Sprintf("%s.md", t.Name)))
		if err != nil {
			_ = file.Close()
			return errors.WithStack(err)
		}

		er := false
		if _, err := os.Lstat(filepath.Join(fullPath, fmt.Sprintf("%s.%s", t.Name, c.ER.Format))); err == nil {
			er = true
		}

		md := New(c, er)

		err = md.OutputTable(file, t)
		if err != nil {
			_ = file.Close()
			return errors.WithStack(err)
		}
		fmt.Printf("%s\n", filepath.Join(docPath, fmt.Sprintf("%s.md", t.Name)))
		err = file.Close()
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// Diff database and markdown files.
func Diff(s *schema.Schema, c *config.Config) (string, error) {
	docPath := c.DocPath

	var diff string
	fullPath, err := filepath.Abs(docPath)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if !outputExists(s, fullPath) {
		return "", errors.New("target files does not exists")
	}

	// README.md
	b := new(bytes.Buffer)
	er := false
	if _, err := os.Lstat(filepath.Join(fullPath, fmt.Sprintf("schema.%s", c.ER.Format))); err == nil {
		er = true
	}

	md := New(c, er)

	_, err = md.OutputSchema(b, s)
	if err != nil {
		return "", errors.WithStack(err)
	}

	targetPath := filepath.Join(fullPath, "README.md")
	a, err := ioutil.ReadFile(filepath.Clean(targetPath))
	if err != nil {
		a = []byte{}
	}

	mdsn, err := c.MaskedDSN()
	if err != nil {
		return "", errors.WithStack(err)
	}
	to := fmt.Sprintf("tbls doc %s", mdsn)

	from := filepath.Join(docPath, "README.md")

	d := difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(a)),
		B:        difflib.SplitLines(b.String()),
		FromFile: from,
		ToFile:   to,
		Context:  3,
	}

	text, _ := difflib.GetUnifiedDiffString(d)
	if text != "" {
		diff += fmt.Sprintf("diff %s '%s'\n", from, to)
		diff += text
	}

	// tables
	for _, t := range s.Tables {
		b := new(bytes.Buffer)
		er := false
		if _, err := os.Lstat(filepath.Join(fullPath, fmt.Sprintf("%s.%s", t.Name, c.ER.Format))); err == nil {
			er = true
		}

		md := New(c, er)

		err := md.OutputTable(b, t)
		if err != nil {
			return "", errors.WithStack(err)
		}
		targetPath := filepath.Join(fullPath, fmt.Sprintf("%s.md", t.Name))
		a, err := ioutil.ReadFile(filepath.Clean(targetPath))
		if err != nil {
			a = []byte{}
		}

		from := filepath.Join(docPath, fmt.Sprintf("%s.md", t.Name))

		d := difflib.UnifiedDiff{
			A:        difflib.SplitLines(string(a)),
			B:        difflib.SplitLines(b.String()),
			FromFile: from,
			ToFile:   to,
			Context:  3,
		}

		text, _ := difflib.GetUnifiedDiffString(d)
		if text != "" {
			diff += fmt.Sprintf("diff %s '%s'\n", from, to)
			diff += text
		}
	}
	return diff, nil
}

func outputExists(s *schema.Schema, path string) bool {
	// README.md
	if _, err := os.Lstat(filepath.Join(path, "README.md")); err == nil {
		return true
	}
	// tables
	for _, t := range s.Tables {
		if _, err := os.Lstat(filepath.Join(path, fmt.Sprintf("%s.md", t.Name))); err == nil {
			return true
		}
	}
	return false
}

func (m *Md) makeSchemaTemplateData(s *schema.Schema, adjust bool) map[string]interface{} {
	tablesData := [][]string{
		[]string{
			m.config.MergedDict.Lookup("Name"),
			m.config.MergedDict.Lookup("Columns"),
			m.config.MergedDict.Lookup("Comment"),
			m.config.MergedDict.Lookup("Type"),
		},
		[]string{"----", "-------", "-------", "----"},
	}
	for _, t := range s.Tables {
		data := []string{
			fmt.Sprintf("[%s](%s.md)", t.Name, t.Name),
			fmt.Sprintf("%d", len(t.Columns)),
			t.Comment,
			t.Type,
		}
		tablesData = append(tablesData, data)
	}

	if adjust {
		return map[string]interface{}{
			"Schema": s,
			"Tables": adjustTable(tablesData),
		}
	}

	return map[string]interface{}{
		"Schema": s,
		"Tables": tablesData,
	}
}

func (m *Md) makeTableTemplateData(t *schema.Table, adjust bool) map[string]interface{} {
	// Columns
	columnsData := [][]string{
		[]string{
			m.config.MergedDict.Lookup("Name"),
			m.config.MergedDict.Lookup("Type"),
			m.config.MergedDict.Lookup("Default"),
			m.config.MergedDict.Lookup("Nullable"),
			m.config.MergedDict.Lookup("Children"),
			m.config.MergedDict.Lookup("Parents"),
			m.config.MergedDict.Lookup("Comment"),
		},
		[]string{"----", "----", "-------", "--------", "--------", "-------", "-------"},
	}
	for _, c := range t.Columns {
		childRelations := []string{}
		cEncountered := map[string]bool{}
		for _, r := range c.ChildRelations {
			if _, ok := cEncountered[r.Table.Name]; ok {
				continue
			}
			childRelations = append(childRelations, fmt.Sprintf("[%s](%s.md)", r.Table.Name, r.Table.Name))
			cEncountered[r.Table.Name] = true
		}
		parentRelations := []string{}
		pEncountered := map[string]bool{}
		for _, r := range c.ParentRelations {
			if _, ok := pEncountered[r.ParentTable.Name]; ok {
				continue
			}
			parentRelations = append(parentRelations, fmt.Sprintf("[%s](%s.md)", r.ParentTable.Name, r.ParentTable.Name))
			pEncountered[r.ParentTable.Name] = true
		}
		data := []string{
			c.Name,
			c.Type,
			c.Default.String,
			fmt.Sprintf("%v", c.Nullable),
			strings.Join(childRelations, " "),
			strings.Join(parentRelations, " "),
			c.Comment,
		}
		columnsData = append(columnsData, data)
	}

	// Constraints
	constraintsData := [][]string{
		[]string{
			m.config.MergedDict.Lookup("Name"),
			m.config.MergedDict.Lookup("Type"),
			m.config.MergedDict.Lookup("Definition"),
		},
		[]string{"----", "----", "----------"},
	}
	cComment := false
	for _, c := range t.Constraints {
		if c.Comment != "" {
			cComment = true
		}
	}
	if cComment {
		constraintsData[0] = append(constraintsData[0], m.config.MergedDict.Lookup("Comment"))
		constraintsData[1] = append(constraintsData[1], "-------")
	}
	for _, c := range t.Constraints {
		data := []string{
			c.Name,
			c.Type,
			c.Def,
		}
		if cComment {
			data = append(data, c.Comment)
		}
		constraintsData = append(constraintsData, data)
	}

	// Indexes
	indexesData := [][]string{
		[]string{
			m.config.MergedDict.Lookup("Name"),
			m.config.MergedDict.Lookup("Definition"),
		},
		[]string{"----", "----------"},
	}
	iComment := false
	for _, i := range t.Indexes {
		if i.Comment != "" {
			iComment = true
		}
	}
	if iComment {
		indexesData[0] = append(indexesData[0], m.config.MergedDict.Lookup("Comment"))
		indexesData[1] = append(indexesData[1], "-------")
	}
	for _, i := range t.Indexes {
		data := []string{
			i.Name,
			i.Def,
		}
		if iComment {
			data = append(data, i.Comment)
		}
		indexesData = append(indexesData, data)
	}

	// Triggers
	triggersData := [][]string{
		[]string{
			m.config.MergedDict.Lookup("Name"),
			m.config.MergedDict.Lookup("Definition"),
		},
		[]string{"----", "----------"},
	}
	tComment := false
	for _, t := range t.Triggers {
		if t.Comment != "" {
			tComment = true
		}
	}
	if tComment {
		triggersData[0] = append(triggersData[0], m.config.MergedDict.Lookup("Comment"))
		triggersData[1] = append(triggersData[1], "-------")
	}
	for _, t := range t.Triggers {
		data := []string{
			t.Name,
			t.Def,
		}
		if tComment {
			data = append(data, t.Comment)
		}
		triggersData = append(triggersData, data)
	}

	if adjust {
		return map[string]interface{}{
			"Table":       t,
			"Columns":     adjustTable(columnsData),
			"Constraints": adjustTable(constraintsData),
			"Indexes":     adjustTable(indexesData),
			"Triggers":    adjustTable(triggersData),
		}
	}

	return map[string]interface{}{
		"Table":       t,
		"Columns":     columnsData,
		"Constraints": constraintsData,
		"Indexes":     indexesData,
		"Triggers":    triggersData,
	}
}

func adjustTable(data [][]string) [][]string {
	r := strings.NewReplacer("\r\n", "<br>", "\n", "<br>", "\r", "<br>")
	w := make([]int, len(data[0]))
	for i := range data {
		for j := range w {
			l := runewidth.StringWidth(r.Replace(data[i][j]))
			if l > w[j] {
				w[j] = l
			}
		}
	}
	for i := range data {
		for j := range w {
			if i == 1 {
				data[i][j] = strings.Repeat("-", w[j])
			} else {
				data[i][j] = fmt.Sprintf(fmt.Sprintf("%%-%ds", w[j]), r.Replace(data[i][j]))
			}
		}
	}

	return data
}
