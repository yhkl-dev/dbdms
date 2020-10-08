package v1

import (
	"dbdms/apps/user"

	"github.com/gin-gonic/gin"
)

// RegisterRouter for apps
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	user.RegisterRouter(v1.Group("/users"))
}
