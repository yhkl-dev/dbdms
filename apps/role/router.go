package role

import "github.com/gin-gonic/gin"

// RegisterRouter role router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", GetAllRoles)
	router.GET("/:id", GetRoleDetail)
	router.PUT("/:id", SaveOrUpdateRole)
	router.POST("", SaveOrUpdateRole)
	router.DELETE("/:id", DeleteRole)
}
