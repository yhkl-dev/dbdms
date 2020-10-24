package rbac

import (
	"dbdms/db"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// E Casbin enforcer
var E *casbin.Enforcer

func init() {
	adapter, err := gormadapter.NewAdapterByDB(db.SQL)

	if err != nil {
		log.Fatal(err)
	}
	e, err := casbin.NewEnforcer("config/model.conf", adapter)
	if err != nil {
		log.Fatal(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal(err)
	}
	E = e

	userRoles := getUserRoles()
	for _, ur := range userRoles {
		_, err := E.AddRoleForUser(ur["user_name"].(string), ur["role_name"].(string))
		if err != nil {
			log.Fatal(err)
		}
	}
	routeRoles := getRouteRoles()
	for _, rr := range routeRoles {
		_, err := E.AddPolicy(rr["role_name"].(string), rr["route_path"].(string), rr["route_method"].(string))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getUserRoles() []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	db.SQL.Select("a.user_name, c.role_name").
		Table("sys_users a, user_role_mapping b, sys_roles c ").
		Where("a.user_id = b.user_id AND b.role_id = c.role_id").
		Order("a.user_id desc").Find(&result)
	return result
}

func getRouteRoles() []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	db.SQL.Select("b.route_path,b.route_method,a.role_name").
		Table("sys_roles a,sys_routes b,role_route_mapping c").
		Where("a.role_id = c.role_id AND c.route_id = b.route_id").
		Order("a.role_id desc").Find(&result)
	return result
}
