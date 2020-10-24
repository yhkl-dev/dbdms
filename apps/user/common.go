package user

import (
	"dbdms/midware/jwtauth"
	"dbdms/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// generateToken
func generateToken(context *gin.Context, user *User) {
	j := jwtauth.NewJWT()
	claims := jwtauth.CustomClaims{
		UserID:    strconv.Itoa(user.UserID),
		UserName:  user.UserName,
		UserPhone: user.UserPhone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() + utils.GetTokenConfig().ActiveTime),       // effective time
			ExpiresAt: int64(time.Now().Unix() + utils.GetTokenConfig().ExpiredTime*3600), // expire time
			Issuer:    utils.GetTokenConfig().Issuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: err.Error(),
		})
		context.Abort()
	}
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Message: utils.StatusText(utils.LoginStatusOK),
		// Content: gin.H{"ACCESS_TOKEN": token, "User": user},
		Content: gin.H{"ACCESS_TOKEN": token},
	})
}
