package rbac

import (
	"dbdms/midware/jwtauth"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

// RBAC midware
func RBAC() gin.HandlerFunc {

	return func(context *gin.Context) {
		fmt.Println("RBAC midware")

		claims, _ := context.Get("claims")
		fmt.Println(reflect.TypeOf(claims))
		fmt.Println("claims.(jwtauth.CustomClaims).UserName", claims.(*jwtauth.CustomClaims).UserName)
		access, err := E.Enforce(claims.(*jwtauth.CustomClaims).UserName, context.Request.RequestURI, context.Request.Method)

		if err != nil || !access {
			userRoles := getUserRoles()
			fmt.Println(userRoles)
			// TODO
			// context.JSON(http.StatusUnprocessableEntity, utils.JSONObject{
			// 	Code:    "0",
			// 	Message: utils.StatusText(utils.BindModelError),
			// 	Content: err,
			// })
			context.Next()
		} else {
			context.Next()
		}
	}
}
