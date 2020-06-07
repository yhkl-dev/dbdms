package user

import (
	helper "dbdms/helpers"
	"dbdms/system"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login login interface
func Login(context *gin.Context) {
	params := &helper.LoginParams{}
	if err := context.Bind(params); err == nil {
		userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
		user := userService.GetUserByName(params.Username)
		if user != nil && user.Password == helper.SHA256(params.Password) {
			user.LoginCount += 1
			user.LoginTime = time.Now()
			err := userService.SaveOrUpdate(user)
			if err == nil {
				generateToken(context, user)
			} else {
				context.JSON(http.StatusOK, helper.JSONObject{
					Code:    "0",
					Message: helper.StatusText(helper.LoginStatusSQLError),
					Content: err,
				})
			}
		} else {
			context.JSON(http.StatusOK, helper.JSONObject{
				Code:    "0",
				Message: helper.StatusText(helper.LoginStatusError),
			})
		}
	} else {
		context.JSON(http.StatusUnprocessableEntity, helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.BindModelError),
			Content: err,
		})
	}
}

// generateToken
func generateToken(context *gin.Context, user *User) {
	j := system.NewJWT()
	claims := system.CustomClaims{
		ID:    string(user.ID),
		Name:  user.UserName,
		Phone: user.Phone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() + system.GetTokenConfig().ActiveTime),       // effective time
			ExpiresAt: int64(time.Now().Unix() + system.GetTokenConfig().ExpiredTime*3600), // expire time
			Issuer:    system.GetTokenConfig().Issuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "0",
			Message: err.Error(),
		})
		context.Abort()

	}
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "1",
		Message: helper.StatusText(helper.LoginStatusOK),
		Content: gin.H{"ACCESS_TOKEN": token, "User": user},
	})

}

func init() {
	// 先读取Token配置文件
	err := system.LoadTokenConfig("./conf/token-config.yml")
	if err != nil {
		helper.ErrorLogger.Errorln("Read Token Config Error: ", err)
	}
	if len(strings.TrimSpace(system.GetTokenConfig().SignKey)) > 0 {
		system.SetSignKey(system.GetTokenConfig().SignKey)
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
