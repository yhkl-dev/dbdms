package permission

import (
	helper "dbdms/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllPermissions(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	pageSize, _ := strconv.Atoi(context.Query("page_size"))
	codeName := context.Query("code_name")
	permissionService := PermissionServiceInstance(PermissionRepositoryInterface(helper.SQL))
	pageBean := permissionService.GetPage(page, pageSize, &Permission{CodeName: codeName})
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
}
