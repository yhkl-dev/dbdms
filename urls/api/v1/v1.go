package v1

import (
	"dbdms/apps/document"
	"dbdms/apps/resources"
	"dbdms/apps/role"
	"dbdms/apps/routes"
	"dbdms/apps/user"

	"github.com/gin-gonic/gin"
)

// RegisterRouter for apps
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	user.RegisterRouter(v1.Group("/users"))
	role.RegisterRouter(v1.Group("/roles"))
	routes.RegisterRouter(v1.Group("/routes"))
	resources.RegisterRouter(v1.Group("/resources"))
	resources.RegisterResourceTypeRouter(v1.Group("/resource_types"))
	document.RegisterRouter(v1.Group("/documents"))
	document.RegisterVersionRouter(v1.Group("/document_versions"))
}
