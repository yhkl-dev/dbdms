package role

import (
	helper "dbdms/helpers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	fmt.Println(roleID)
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
