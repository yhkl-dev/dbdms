package permission

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllPermissions)
}
