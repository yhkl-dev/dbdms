package urls

import (
	"dbdms/apps/user"
	"dbdms/midware/jwtauth"
	"dbdms/midware/rbac"
	v1 "dbdms/urls/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RegisterAPIRoutes register all api routers
func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("api")
	api.Use(jwtauth.AUTH())
	api.Use(rbac.RBAC())
	{
		v1.RegisterRouter(api)
	}
}

// RegisterOpenRoutes register routes which does not need auth
func RegisterOpenRoutes(router *gin.Engine) {
	router.POST("login", user.Login)
	router.POST("register", user.Register)
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
