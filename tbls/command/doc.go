package command

import (
	"dbdms/tbls/config"
	"dbdms/tbls/schema"
	"dbdms/tbls/output/gviz"
	"dbdms/tbls/output/md"
	"dbdms/tbls/datasource"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
)

func Doc(args []string) {
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	s, err := datasource.Analyze(c.DSN)
	if err != nil {
		log.Fatal(err)
	}

	err = c.ModifySchema(s)
	if err != nil {
		log.Fatal(err)
	}

	err = md.Output(s, c, false)

	if err != nil {
		log.Fatal(err)
	}
}

func withDot(s *schema.Schema, c *config.Config, force bool) (e error) {
	erFormat := c.ER.Format
	outputPath := c.DocPath
	fullPath, err := filepath.Abs(outputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	err = os.MkdirAll(fullPath, 0755) // #nosec
	if err != nil {
		return errors.WithStack(err)
	}

	erFileName := fmt.Sprintf("schema.%s", erFormat)
	fmt.Printf("%s\n", filepath.Join(outputPath, erFileName))

	file, err := os.OpenFile(filepath.Join(fullPath, erFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // #nosec
	if err != nil {
		return errors.WithStack(err)
	}
	g := gviz.New(c)
	err = g.OutputSchema(file, s)
	if err != nil {
		return errors.WithStack(err)
	}

	// tables
	for _, t := range s.Tables {
		erFileName := fmt.Sprintf("%s.%s", t.Name, erFormat)
		fmt.Printf("%s\n", filepath.Join(outputPath, erFileName))

		file, err := os.OpenFile(filepath.Join(fullPath, erFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // #nosec
		if err != nil {
			return errors.WithStack(err)
		}
		err = g.OutputTable(file, t)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func loadDocArgs(args []string) ([]config.Option, error) {
	options := []config.Option{}
	if len(args) > 2 {
		return options, errors.WithStack(errors.New("too many arguments"))
	}
	if len(args) == 2 {
		options = append(options, config.DSNURL(args[0]))
		options = append(options, config.DocPath(args[1]))
	}
	if len(args) == 1 {
		options = append(options, config.DSNURL(args[0]))
	}
	return options, nil
}
