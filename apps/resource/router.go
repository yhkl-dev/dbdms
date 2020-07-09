package resource

import "github.com/gin-gonic/gin"

// RegisterRouter use router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", GetAllDBResources)
	router.GET("/:id", GetDBResource)
	router.DELETE("/:id", DeleteDBResource)
	router.PUT("/:id", UpdateDBResource)
}
