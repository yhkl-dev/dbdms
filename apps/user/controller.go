package user

import (
	"dbdms/db"
	"dbdms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAllUsers 获取用户列表
// @Summary 获取用户列表
// @Tags UserController
// @Accept json
// @Produce json
// @Success 200 {stirng} json {}
// @Router /api/v1/users [get]
func ListAllUsers(context *gin.Context) {
	query := utils.UserQueryParams{}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.ParamParseError),
			Content: err.Error(),
		})
	}
	userService := ServiceInstance(RepoInterface(db.SQL))
	pageBean := userService.GetPage(query.Page, query.PageSize, &User{UserName: query.UserName, UserPhone: query.UserPhone, UserEmail: query.UserEmail})
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}

// Register 用户注册
// @summary 用户注册方法
// @Tags UserController
// @Accept json
func Register(context *gin.Context) {
	user := &User{}
	err := context.Bind(user)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err.Error(),
		})
	}
	err = user.validator()
	if err != nil {
		context.JSON(http.StatusOK, &utils.JSONObject{
			Code:    "0",
			Message: err.Error(),
		})
		return
	}
	userService := ServiceInstance(RepoInterface(db.SQL))
	err = userService.SaveOrUpdate(user)
	if err == nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.SaveStatusOK),
		})
		return
	}
	context.JSON(http.StatusOK, &utils.JSONObject{
		Code:    "1",
		Message: err.Error(),
	})
}

// Login user login method
func Login(context *gin.Context) {
	params := &utils.LoginParams{}
	if err := context.Bind(params); err == nil {
		userService := ServiceInstance(RepoInterface(db.SQL))
		user := userService.GetUserByName(params.UserName)
		if user != nil && user.UserPassword == utils.SHA256(params.UserPassword) {
			err := userService.SaveOrUpdate(user)
			if err != nil {
				context.JSON(http.StatusOK, utils.JSONObject{
					Code:    "0",
					Message: utils.StatusText(utils.LoginStatusSQLError),
					Content: err,
				})
			}
			generateToken(context, user)
		} else {
			context.JSON(http.StatusOK, utils.JSONObject{
				Code:    "0",
				Content: "",
				Message: utils.StatusText(utils.LoginStatusError),
			})
		}
	} else {
		context.JSON(http.StatusUnprocessableEntity, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err,
		})
	}
}
