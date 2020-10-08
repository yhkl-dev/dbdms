package rbac

import (
	"dbdms/db"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// E casbin enforcer
var E *casbin.Enforcer

func init() {
	adapter, err := gormadapter.NewAdapterByDB(db.SQL)

	if err != nil {
		log.Fatal()
	}
	e, err := casbin.NewEnforcer("config/model.conf", adapter)
	if err != nil {
		log.Fatal()
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal()
	}
	E = e
}
