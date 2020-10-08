package role

import "github.com/gin-gonic/gin"

// RegisterRouter use router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllRoles)
	router.POST("", AddRole)
	// router.DELETE("/:id", DeleteUser)
	// router.PUT("/:id", UpdateUserProfile)
}
