package document

import "github.com/gin-gonic/gin"

// RegisterRouter resource router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllResourcesDocuments)
}