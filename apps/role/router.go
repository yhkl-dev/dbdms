package role

import "github.com/gin-gonic/gin"

// RegisterRouter use router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllRoles)
	router.POST("", AddRole)
	router.POST("/user/:id", ChangeUserRole)
	router.DELETE("/:id", DeleteRoleByID)
	router.PUT("/:id", UpdateRole)
}
