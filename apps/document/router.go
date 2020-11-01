package document

import "github.com/gin-gonic/gin"

// RegisterRouter resource router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", ListAllResourcesDocuments)
	router.GET("/resource/:id", GenerateDocument)
}

// RegisterRouter resource router register
func RegisterVersionRouter(router *gin.RouterGroup) {
	router.GET("", ListDocumentVersions)
}