package document

import (
	"dbdms/tbls/config"
	"dbdms/tbls/datasource"
	"dbdms/tbls/output/md"
	"dbdms/tbls/schema"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type documentQueryParams struct {
	DocumentDBName    string `json:"document_db_name" form:"resource_name"`
	DocumentTableName string `json:"document_table_name" form:"resource_host_ip"`
	ResourceName      string `json:"resource_name" form:"resource_name"`
	Page              int    `json:"page" form:"page"`
	PageSize          int    `json:"page_size" form:"page_size"`
}

func Doc(dsn string, resourceID int, documentService Service) {
	args := []string{dsn}
	fmt.Println("start")
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	configPath := "./docs/.tbls.yml"

	options, err := loadDocArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Load(configPath, options...)
	if err != nil {
		log.Fatal(err)
	}

	s, err := datasource.Analyze(c.DSN)
	fmt.Println(s)
	if err != nil {
		log.Fatal(err)
	}

	err = c.ModifySchema(s)
	if err != nil {
		log.Fatal(err)
	}

	err = Output(s, c, resourceID, documentService)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("end")
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

// Output generate markdown files.
func Output(s *schema.Schema, c *config.Config, resourceID int, documentService Service) (e error) {
	document := &DatabaseDocument{}
	docPath := c.DocPath

	fullPath, err := filepath.Abs(docPath)
	if err != nil {
		return errors.WithStack(err)
	}
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
	// README.md
	mdPointer := md.New(c, false)
	// TODO 现在的逻辑是读取文件里面的内容，后续需要优化直接读取io.Writer里面的数据流
	_, err = mdPointer.OutputSchema(file, s)
	if err != nil {
		return errors.WithStack(err)
	}
	bytes, err := ioutil.ReadFile(filepath.Join(fullPath, "README.md"))
	if err != nil {
		fmt.Println("error : %s", err)
		return
	}
	document.DocumentContent = string(bytes)
	document.ResourceID = resourceID
	document.DocumentDBName = s.Name
	document.DocumentTableName = "README"
	document.DocumentFileName = "README.md"
	err = documentService.SaveDocument(document)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("README.md")
	//fmt.Printf("%s\n", filepath.Join(docPath, "README.md"))

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

		mdPointer := md.New(c, er)

		err = mdPointer.OutputTable(file, t)
		if err != nil {
			_ = file.Close()
			return errors.WithStack(err)
		}
		// ---------------------
		bytes, err := ioutil.ReadFile(filepath.Join(fullPath, fmt.Sprintf("%s.md", t.Name)))
		if err != nil {
			fmt.Println("error : %s", err)
			return
		}
		documentTable := new(DatabaseDocument)
		documentTable.DocumentContent = string(bytes)
		documentTable.ResourceID = resourceID
		documentTable.DocumentDBName = s.Name
		documentTable.DocumentTableName = t.Name
		documentTable.DocumentFileName = fmt.Sprintf("%s.md", t.Name)
		err = documentService.SaveDocument(documentTable)
		if err != nil {
			fmt.Println(err)
		}

		// ---------------------
		err = file.Close()
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
