package routes

import (
	"dbdms/db"
	"dbdms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAllRoutes 获取路由列表
func ListAllRoutes(context *gin.Context) {
	query := roleQueryParams{}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.ParamParseError),
			Content: err.Error(),
		})
	}
	routeService := ServiceInstance(RepoInterface(db.SQL))
	// pageBean := routeService.GetPage(query.Page, query.PageSize, &Routes{})
	pageBean := routeService.ListAllRoutes()
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}
