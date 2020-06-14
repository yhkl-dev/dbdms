package role

import (
	helper "dbdms/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SaveRole(context *gin.Context) {
	role := &Role{}
	err := context.Bind(&role)
	if err == nil {
		role.DeleteAt = nil
		roleService := RoleServiceInstance(RoleRepositoryIntance(helper.SQL))
		err := roleService.SaveOrUpdate(role)
		if err == nil {
			context.JSON(http.StatusOK, &helper.JSONObject{
				Code:    "1",
				Message: helper.StatusText(helper.SaveStatusOK),
			})
			return
		}
		context.JSON(http.StatusOK, &helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.SaveStatusError),
			Content: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "0",
		Message: helper.StatusText(helper.BindModelError),
		Content: err,
	})
	return
}

func SaveOrUpdateRole(context *gin.Context) {
	roleIDString := context.Param("id")
	roleID, err := strconv.Atoi(roleIDString)
	if err != nil {
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.ParamParseError),
			Content: err,
		})
		return
	}
	var role Role
	role.ID = roleID
	if err := context.Bind(&role); err == nil {
		role.DeleteAt = nil
		roleService := RoleServiceInstance(RoleRepositoryIntance(helper.SQL))
		err := roleService.SaveOrUpdate(&role)
		if err == nil {
			context.JSON(http.StatusOK, &helper.JSONObject{
				Code:    "1",
				Message: helper.StatusText(helper.SaveStatusOK),
			})
			return
		}
		context.JSON(http.StatusOK, &helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.SaveStatusError),
			Content: err.Error(),
		})
		return
	} else {
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.BindModelError),
			Content: err,
		})
		return
	}
}

func GetAllRoles(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	pageSize, _ := strconv.Atoi(context.Query("page_size"))
	rolename := context.Query("rolename")
	roleService := RoleServiceInstance(RoleRepositoryIntance(helper.SQL))
	pageBean := roleService.GetPage(page, pageSize, &Role{RoleName: rolename})
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "1",
		Content: pageBean,
	})

}

func GetRoleDetail(context *gin.Context) {
	roleIDString := context.Param("id")
	roleService := RoleServiceInstance(RoleRepositoryIntance(helper.SQL))
	roleID, err := strconv.Atoi(roleIDString)
	if err == nil {
		role := roleService.GetByID(roleID)
		if role != nil {
			context.JSON(http.StatusOK, helper.JSONObject{
				Code:    "1",
				Content: role,
			})
			return
		}
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.ResourceDoesNotExist),
		})
		return
	}
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "0",
		Message: helper.StatusText(helper.ParamParseError),
		Content: err.Error(),
	})
	return
}

func DeleteRole(context *gin.Context) {
	roleIDString := context.Param("id")
	roleService := RoleServiceInstance(RoleRepositoryIntance(helper.SQL))
	roleID, err := strconv.Atoi(roleIDString)
	if err == nil {
		err = roleService.DeleteByID(roleID)
		if err == nil {
			context.JSON(http.StatusOK, helper.JSONObject{
				Code:    "1",
				Message: helper.StatusText(helper.DeleteStatusOK),
			})
			return

		}
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "0",
			Message: err.Error(),
		})
		return

	}
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "0",
		Message: helper.StatusText(helper.ParamParseError),
		Content: err.Error(),
	})
	return
}
