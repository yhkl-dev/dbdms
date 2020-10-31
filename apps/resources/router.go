package resources


import "github.com/gin-gonic/gin"

// RegisterRouter resource router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllResources)
	router.POST("", CreateResource)
	router.PUT("/:id", UpdateResource)
	router.DELETE("/:id", DeleteResourceByID)
}

// RegisterResourceTypeRouter resource type router register
func RegisterResourceTypeRouter(router *gin.RouterGroup) {
	router.GET("", ListAllResourceTypes)
	router.POST("", CreateResourceType)
	router.PUT("/:id", UpdateResourceType)
	router.DELETE("/:id", DeleteResourceTypeByID)
}