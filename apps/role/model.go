package role

import (
	"dbdms/db"
	"fmt"
)

// Role role struct
type Role struct {
	RoleID      int    `gorm:"column:role_id;primary_key"`
	RoleName    string `gorm:"type:varchar(32);column:role_name;unique;not null" json:"role_name" form:"role_name" binding:"required"`
	RolePID     int    `gorm:"column:role_pid" json:"role_pid" form:"role_pid"`
	RoleComment string `gorm:"type:varchar(32);column:role_comment" json:"role_comment" form:"role_comment" binding:"required"`
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

// Validator role validator
func (role *Role) Validator() error {
	// if ok, err := regex.MatchLetterNumMinAndMax(role.RoleName, 4, 6, "role_name"); !ok {
	// if role.RoleName == "" {
	// 	return errors.New("role name can not be null ")
	// }
	// if role.RoleComment == "" {
	// return errors.New("role description can not be null ")
	// }
	return nil
}

// TableName define table name
func (urm *UserRoleMapping) TableName() string {
	return "roles_users_mapping"
}

func (role *Role) String() string {
	return fmt.Sprintf("<RoleID: %d, RoleName: %s, RolePID: %d, RoleDescription: %s>", role.RoleID, role.RoleName, role.RolePID, role.RoleComment)
}

func init() {
	db.SQL.AutoMigrate(&Role{})
	db.SQL.AutoMigrate(&UserRoleMapping{})
}
