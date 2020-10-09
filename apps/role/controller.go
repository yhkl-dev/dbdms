package role

import (
	"dbdms/db"
	"dbdms/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListAllRoles 获取角色列表
func ListAllRoles(context *gin.Context) {
	query := roleQueryParams{}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.ParamParseError),
			Content: err.Error(),
		})
	}
	roleService := ServiceInstance(RepoInterface(db.SQL))
	pageBean := roleService.GetPage(query.Page, query.PageSize, &Role{RoleName: query.RoleName})
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}

// AddRole add role 添加角色
func AddRole(context *gin.Context) {
	role := &Role{}
	err := context.Bind(role)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err.Error(),
		})
		return
	}
	roleService := ServiceInstance(RepoInterface(db.SQL))
	err = roleService.SaveOrUpdate(role)
	if err == nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.SaveStatusOK),
		})
		return
	}
	context.JSON(http.StatusOK, &utils.JSONObject{
		Code:    "0",
		Message: err.Error(),
	})
}

// DeleteRoleByID delete role 删除角色
func DeleteRoleByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	roleService := ServiceInstance(RepoInterface(db.SQL))
	err := roleService.DeleteRoleByID(id)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.DeleteStatusErr),
			Content: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Message: utils.StatusText(utils.DeleteStatusOK),
	})
}

// UpdateRole update role info 更新角色信息
func UpdateRole(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	role := &Role{}
	err := context.Bind(role)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err.Error(),
		})
		return
	}
	roleService := ServiceInstance(RepoInterface(db.SQL))
	x := roleService.GetByID(id)
	if x == nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.UpdateObjIsNil),
		})
		return
	}
	role.RoleID = id

	err = roleService.SaveOrUpdate(role)
	if err == nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.SaveStatusOK),
		})
		return
	}
	context.JSON(http.StatusOK, &utils.JSONObject{
		Code:    "0",
		Message: err.Error(),
	})
}
