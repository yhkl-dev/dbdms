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
	//userPerm := user.User{}.RegisterPermission()
	//for _, pn := range userPerm {
	//	fmt.Println(pn)
	//}
	//	helper.SQL.Create()
	var permList = make(map[string]map[string]string)
	permList["can_view_user"] = map[string]string{"id": "1", "ModelName": "user", "PermissionName": "GET", "CodeName": "/api/v1/user"}
	permList["can_add_user"] = map[string]string{"id": "2", "ModelName": "user", "PermissionName": "POST", "CodeName": "/api/v1/user/(.id+)"}
	permList["can_update_user"] = map[string]string{"id": "3", "ModelName": "user", "PermissionName": "PUT", "CodeName": "/api/v1/user/(.d+)"}
	permList["can_delete_user"] = map[string]string{"id": "4", "ModelName": "user", "PermissionName": "DELETE", "CodeName": "/api/v1/user/(.d+)"}
	for cn, pn := range permList {
		id, _ := strconv.Atoi(pn["id"])
		sql := fmt.Sprintf("INSERT IGNORE INTO permission (id,  mode_name, permission_name, code_name, description ) VALUES (%v, \"%v\",\"%v\", \"%v\", \"%v\" )", id, pn["ModelName"], pn["PermissionName"], pn["CodeName"], cn)
		helper.SQL.Exec(sql)
	}
}
