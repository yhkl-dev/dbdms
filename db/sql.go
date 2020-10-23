package db

import (
	"dbdms/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// SQL database Operator
var SQL *gorm.DB

func init() {
	err := utils.LoadDatasourceConfig("./config/datasource.yml")

	if err != nil {
		// ErrorLogger.Errorln("Read Database Config Failed: ", err)
		log.Fatal(err)
	}
	datasource := utils.GetDatasource()
	SQL, err = gorm.Open(mysql.Open(datasource.Username+":"+datasource.Password+datasource.URL), &gorm.Config{})
	if err != nil {
		// ErrorLogger.Errorln("Connect Database Error: ", err)
		log.Fatal(err)
	}
	db, err := SQL.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(datasource.MaxOpenConns)
	db.SetMaxIdleConns(datasource.MaxIdleConns)
}
