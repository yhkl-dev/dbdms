package permission

import (
	helper "dbdms/helpers"
)

// Permission for permission table
type Permission struct {
	ID             int    `gorm:"primary_key;column:id"`
	PermissionName string `gorm:"type:varchar(32);not null;column:permission_name"`
	CodeName       string `gorm:"type:varchar(32);not null;column:code_name"`
	Description    string `gorm:"type:varchar(100);not null;column:description"`
}

func init() {
	helper.SQL.AutoMigrate(&Permission{})
}
