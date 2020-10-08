package routes

import (
	"dbdms/db"
	"fmt"
)

// Role role struct
type Role struct {
	RoleID      int    `gorm:"column:role_id; primary_key"`
	RoleName    string `gorm:"column:role_name"`
	RolePID     int    `gorm:"column:role_pid"`
	RoleComment string `gorm:"column:role_comment"`
}

// UserRoleMapping user and role mapping table
type UserRoleMapping struct {
	MappingID int `gorm:"column:mapping_id; primary_key"`
	UserID    int `gorm:"column:user_id"`
	RoleID    int `gorm:"column:role_id"`
}

// TableName define table name
func (role *Role) TableName() string {
	return "sys_roles"
}

// TableName define table name
func (urm *UserRoleMapping) TableName() string {
	return "roles_users_mapping"
}

func (role *Role) String() string {
	return fmt.Sprintf("<RoleID: %d, RoleName: %s>", role.RoleID, role.RoleName)
}

func init() {
	db.SQL.AutoMigrate(&Role{})
	db.SQL.AutoMigrate(&UserRoleMapping{})
}
