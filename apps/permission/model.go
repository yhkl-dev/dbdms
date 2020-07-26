package permission

import (
	helper "dbdms/helpers"
	"fmt"
	"strconv"
)

// Permission for permission table
type Permission struct {
	ID             int    `gorm:"primary_key;column:id"`
	ModelName      string `gorm:"type:varchar(32);not null;column:mode_name;json:code_name"`
	PermissionName string `gorm:"type:varchar(32);not null;column:permission_name;json:permission_name"`
	CodeName       string `gorm:"type:varchar(32);not null;column:code_name;json:code_name"`
	Description    string `gorm:"type:varchar(100);default:'-';column:description"`
}

func init() {
	helper.SQL.AutoMigrate(&Permission{})
	var permList = make(map[string]map[string]string)
	permList["can_view_users"] = map[string]string{"id": "1", "ModelName": "user", "PermissionName": "GET", "CodeName": "GET:/api/v1/user"}
	permList["can_view_user"] = map[string]string{"id": "2", "ModelName": "user", "PermissionName": "GET", "CodeName": "GET:/api/v1/user/:id"}
	permList["can_add_user"] = map[string]string{"id": "3", "ModelName": "user", "PermissionName": "POST", "CodeName": "POST:/api/v1/user/:id"}
	permList["can_update_user"] = map[string]string{"id": "4", "ModelName": "user", "PermissionName": "PUT", "CodeName": "PUT:/api/v1/user/:id"}
	permList["can_delete_user"] = map[string]string{"id": "5", "ModelName": "user", "PermissionName": "DELETE", "CodeName": "DELETE:/api/v1/user/:id"}
	permList["can_views_roles"] = map[string]string{"id": "6", "ModelName": "role", "PermissionName": "GET", "CodeName": "GET:/api/v1/role"}
	permList["can_view_role"] = map[string]string{"id": "7", "ModelName": "role", "PermissionName": "GET", "CodeName": "GET:/api/v1/role/:id"}
	permList["can_add_role"] = map[string]string{"id": "8", "ModelName": "role", "PermissionName": "POST", "CodeName": "POST:/api/v1/role/:id"}
	permList["can_update_role"] = map[string]string{"id": "9", "ModelName": "role", "PermissionName": "PUT", "CodeName": "PUT:/api/v1/role/:id"}
	permList["can_delete_role"] = map[string]string{"id": "10", "ModelName": "role", "PermissionName": "DELETE", "CodeName": "DELETE:/api/v1/role/:id"}
	for cn, pn := range permList {
		id, _ := strconv.Atoi(pn["id"])
		sql := fmt.Sprintf("INSERT IGNORE INTO permission (id,  mode_name, permission_name, code_name, description ) VALUES (%v, \"%v\",\"%v\", \"%v\", \"%v\" )", id, pn["ModelName"], pn["PermissionName"], pn["CodeName"], cn)
		helper.SQL.Exec(sql)
	}
}
