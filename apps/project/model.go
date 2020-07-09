package project

import helper "dbdms/helpers"

// Project Struct
type Project struct {
	ID          int    `gorm:"AUTO_INCREMENT;primary_key"`
	ProjectName string `gorm:"type:varchar(32);not null;comment:'项目名称'" json:"project_name" form:"project_name"`
	Description string `gorm:"type:varchar(128);not null;" json:"description" form:"description"`
}

func init() {
	helper.SQL.AutoMigrate(&Project{})
}
