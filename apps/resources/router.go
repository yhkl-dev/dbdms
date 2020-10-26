package resources


import "github.com/gin-gonic/gin"

// RegisterRouter resource router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllResources)
	router.POST("", CreateResource)
	router.PUT("/:id", UpdateResource)
	router.DELETE("/:id", DeleteResourceByID)
}
