package resources


import "github.com/gin-gonic/gin"

// RegisterRouter resource router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllResources)
	router.POST("", CreateResource)
	//router.POST("/user/:id", ChangeUserRole)
	//router.DELETE("/:id", DeleteRoleByID)
	//router.PUT("/:id", UpdateRole)
}
