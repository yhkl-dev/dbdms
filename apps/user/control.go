package user

import (
	helper "dbdms/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	params := &helper.LoginParams{}
	if err := context.Bind(params); err == nil {
		userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
		user := userService.GetUserByName(params.Username)
		if user != nil && user.Password == helper.SHA256(params.Password) {
			user.LoginCount += 1
			user.LoginTiem = time.Now()
			err := userService.SaveOrUpdate(user)
		}
	}
}

// 获取所有用户信息
// @Summary 获取所有用户信息
// @Tags UserController
// @Accept json
// @Produce json
// @Success 200 {object} helpers.JsonObject
// @Router /api/get_all_users [get]
func GetAllUsers(context *gin.Context) {
	userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
	users := userService.GetAll()
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "1",
		Content: users,
	})
	return

}
