package helper

import (
	"dbdms/system"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQL sql handler
var SQL *gorm.DB

func init() {
	err := system.LoadDatasourceConfig("./conf/datasource.yml")

	if err != nil {
		ErrorLogger.Errorln("Read Database Config Failed: ", err)
	}
	datasource := system.GetDatasource()
	SQL, err = gorm.Open(datasource.Driver, datasource.Username+":"+datasource.Password+datasource.URL)
	if err != nil {
		ErrorLogger.Errorln("Connect Database Error: ", err)
		os.Exit(0)
	}
	SQL.DB().SetMaxOpenConns(datasource.MaxOpenConns)
	SQL.DB().SetMaxIdleConns(datasource.MaxIdleConns)
	SQL.SetLogger(SQLLogger)
	SQL.LogMode(datasource.ShowSQL)
	SQL.SingularTable(datasource.SingularTable)
}
