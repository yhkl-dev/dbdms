package routes

import (
	"dbdms/apps"
	"dbdms/utils"

	"gorm.io/gorm"
)

// Repo route interface inplemented from common interface
type Repo interface {
	apps.RepositoryInterface
}

type routeRepo struct {
	db *gorm.DB
}

type roleRouteMappingRepo struct {
	db *gorm.DB
}

var routerRepoInstance = &routeRepo{}
var roleRouteRepoInstance = &roleRouteMappingRepo{}

func RepoInterface(db *gorm.DB) Repo {
	routerRepoInstance.db = db
	roleRouteRepoInstance.db = db
	return routerRepoInstance
}

func (urm *roleRouteMappingRepo) Delete(routeID int) error {
	err := urm.db.Where("route_id = ?", routeID).Delete(RoleRouteMapping{}).Error
	return err
}
func (rp *routeRepo) Insert(m interface{}) error {
	err := rp.db.Create(m).Error
	return err
}

func (rp *routeRepo) Update(m interface{}) error {
	err := rp.db.Save(m).Error
	return err
}

func (rp *routeRepo) Delete(m interface{}) error {
	err := rp.db.Delete(m).Error
	if err != nil {
		return err
	}
	var roleRouteMappingObject roleRouteMappingRepo
	err = roleRouteMappingObject.Delete(m.(Routes).RouteID)
	return err
}

func (rp *routeRepo) FindMore(condition string, params ...interface{}) interface{} {
	routes := make([]*Routes, 0)
	rp.db.Where(condition, params).Find(&routes)
	return routes
}

func (rp *routeRepo) FindOne(id int) interface{} {
	var route Routes
	err := rp.db.Where("route_id = ?", id).First(&route).Error
	if route.RouteID == 0 || err != nil {
		return nil
	}
	return &route
}

func (rp *routeRepo) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean) {
	var total int64
	//	rows := make([]User, 0)
	var rows []Routes

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			rp.db = rp.db.Where(k, v)
		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			rp.db = rp.db.Where(k, v)
		}
	}
	rp.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&rows).Count(&total)
	return &utils.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}

func (rp *routeRepo) FindSingle(condition string, params ...interface{}) interface{} {
	var route Routes
	rp.db.Where(condition, params).First(&route)
	if route.RouteName != "" {
		return &route
	}
	return nil
}
