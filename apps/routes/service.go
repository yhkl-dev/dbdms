package routes

import "dbdms/utils"

type Service interface {
	ListAllRoutes() []*Routes
	GetPage(page int, pageSize int, route *Routes) *utils.PageBean
}

type routeService struct {
	repo Repo
}

var routeServiceIns = &routeService{}

// ServiceInstance 获取route Service 实例
func ServiceInstance(repo Repo) Service {
	routeServiceIns.repo = repo
	return routeServiceIns
}

func (rs *routeService) ListAllRoutes() []*Routes {
	routes := rs.repo.FindMore("1=1").([]*Routes)
	return routes
}

func (rs *routeService) GetPage(page int, pageSize int, route *Routes) *utils.PageBean {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	addCons := make(map[string]interface{})
	if route != nil && route.RouteName != "" {
		addCons["route_name LIKE ?"] = "%" + route.RouteName + "%"
	}

	pageBean := rs.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}
