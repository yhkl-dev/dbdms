package output

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"dbdms/tbls/dict"
	"dbdms/tbls/schema"
)

// Output is interface for output
type Output interface {
	OutputSchema(wr io.Writer, s *schema.Schema) error
	OutputTable(wr io.Writer, s *schema.Table) error
}

func Funcs(d *dict.Dict) map[string]interface{} {
	return template.FuncMap{
		"nl2br": func(text string) string {
			r := strings.NewReplacer("\r\n", "<br>", "\n", "<br>", "\r", "<br>")
			return r.Replace(text)
		},
		"nl2br_slash": func(text string) string {
			r := strings.NewReplacer("\r\n", "<br />", "\n", "<br />", "\r", "<br />")
			return r.Replace(text)
		},
		"nl2mdnl": func(text string) string {
			r := strings.NewReplacer("\r\n", "  \n", "\n", "  \n", "\r", "  \n")
			return r.Replace(text)
		},
		"nl2space": func(text string) string {
			r := strings.NewReplacer("\r\n", " ", "\n", " ", "\r", " ")
			return r.Replace(text)
		},
		"escape_nl": func(text string) string {
			r := strings.NewReplacer("\r\n", "\\n", "\n", "\\n", "\r", "\\n")
			return r.Replace(text)
		},
		"lookup": func(text string) string {
			return d.Lookup(text)
		},
		"label_join": func(labels schema.Labels) string {
			m := []string{}
			for _, l := range labels {
				m = append(m, l.Name)
			}
			return fmt.Sprintf("`%s`", strings.Join(m, "` `"))
		},
	}
}
