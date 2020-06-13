package role

import (
	"dbdms/apps/permission"
	helper "dbdms/helpers"
	"time"
)

// Role role struct
type Role struct {
	ID          int                     `gorm:"primary_key;column:id"`
	RoleName    string                  `gorm:"type:varchar(32);column:role_name;" json:"rolename"`
	Description string                  `gorm:"type:varchar(200);column:description" json:"description"`
	CreateAt    time.Time               `gorm:"column:create_at;default:current_timestamp"`
	UpdateAt    time.Time               `gorm:"column:update_at;default:current_timestamp ON update current_timestamp"`
	IsDeleted   int                     `gorm:"type:int;default:0" json:"is_deleted"` // 0: no 1: yes
	DeleteAt    *time.Time              `gorm:"column:delete_at"`
	Permissions []permission.Permission `gorm:"many2many:role_permission_mapping"`
}

func init() {
	helper.SQL.AutoMigrate(&Role{})
}
