package rbac

import (
	"dbdms/midware/jwtauth"
	"dbdms/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// RBAC middleware
func RBAC() gin.HandlerFunc {

	return func(context *gin.Context) {
		fmt.Println("RBAC midware")

		claims, _ := context.Get("claims")
		uri := strings.TrimPrefix(context.Request.URL.Path, "?")
		access, err := E.Enforce(claims.(*jwtauth.CustomClaims).UserName, uri, context.Request.Method)

		if err != nil || !access {
			//context.AbortWithStatusJSON(403, gin.H{"message": "forbidden"})
			context.AbortWithStatusJSON(http.StatusForbidden, utils.JSONObject{
				Code:    "1",
				Message: utils.StatusText(utils.PermissionDenied),
			})
		} else {
			context.Next()
		}
	}
}
