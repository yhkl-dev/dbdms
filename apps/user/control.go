package user

import (
	helper "dbdms/helpers"
	"dbdms/system"
	"net/http"
	"strconv"
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
			currentTime := time.Now()
			user.LoginTime = &currentTime
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

func Register(context *gin.Context) {
	user := &User{}
	if err := context.Bind(user); err == nil {
		err = user.Validator()
		if err != nil {
			context.JSON(http.StatusOK, &helper.JSONObject{
				Code:    "0",
				Message: err.Error(),
			})
			return
		}
		user.CreateAt = time.Now()
		user.DeleteAt = nil
		user.LoginTime = nil
		userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
		err := userService.SaveOrUpdate(user)
		if err == nil {
			context.JSON(http.StatusOK, helper.JSONObject{
				Code:    "1",
				Message: helper.StatusText(helper.SaveStatusOK),
			})
			return
		}
		context.JSON(http.StatusOK, &helper.JSONObject{
			Code:    "0",
			Message: err.Error(),
		})
		return

	} else {
		context.JSON(http.StatusUnprocessableEntity, &helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.BindModelError),
			Content: err.Error(),
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

// GetAllUsers 获取所有用户信息
// @Summary 获取所有用户信息
// @Tags UserController
// @Accept json
// @Produce json
// @Success 200 {object} helpers.JsonObject
// @Router /api/v1/user [get]
func GetAllUsers(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	pageSize, _ := strconv.Atoi(context.Query("page_size"))
	username := context.Query("username")
	phone := context.Query("phone")
	email := context.Query("email")
	userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
	pageBean := userService.GetPage(page, pageSize, &User{UserName: username, Phone: phone, Email: email})
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
}

func GetUserProfile(context *gin.Context) {
	userIDString := context.Param("id")
	userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
	userID, err := strconv.Atoi(userIDString)
	if err == nil {
		user := userService.GetByID(userID)
		if user != nil {
			context.JSON(http.StatusOK, helper.JSONObject{
				Code:    "1",
				Content: user,
			})
			return
		}
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.UserDoesNotExist),
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

func DeleteUser(context *gin.Context) {
	userIDString := context.Param("id")
	userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
	userID, err := strconv.Atoi(userIDString)
	if err == nil {
		err := userService.DeleteByID(userID)
		if err != nil {
			context.JSON(http.StatusOK, helper.JSONObject{
				Code:    "0",
				Message: helper.StatusText(helper.DeleteStatusErr),
				Content: err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "1",
			Message: helper.StatusText(helper.DeleteStatusOK),
		})
		return
	}
	context.JSON(http.StatusOK, helper.JSONObject{
		Code:    "0",
		Message: helper.StatusText(helper.ParamParseError),
		Content: err,
	})
	return
}

func UpdateUserProfile(context *gin.Context) {
	userIDString := context.Param("id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		context.JSON(http.StatusOK, helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.ParamParseError),
			Content: err,
		})
		return
	}
	user := &User{}
	user.ID = userID
	if err := context.Bind(user); err == nil {
		err = user.Validator()
		if err != nil {
			context.JSON(http.StatusOK, &helper.JSONObject{
				Code:    "0",
				Message: err.Error(),
			})
			return
		}
		userService := UserServiceInstance(UserRepositoryInterface(helper.SQL))
		err := userService.SaveOrUpdate(user)
		if err == nil {
			context.JSON(http.StatusOK, helper.JSONObject{
				Code:    "1",
				Message: helper.StatusText(helper.SaveStatusOK),
			})
			return
		} else {
			context.JSON(http.StatusOK, &helper.JSONObject{
				Code:    "0",
				Message: err.Error(),
			})
			return
		}
	} else {
		context.JSON(http.StatusUnprocessableEntity, helper.JSONObject{
			Code:    "0",
			Message: helper.StatusText(helper.BindModelError),
			Content: err.Error(),
		})
	}
}
