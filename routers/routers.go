package routers

import (
	"dbdms/apps/user"
	v1 "dbdms/routers/api/v1"
	"dbdms/system"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RegisterAPIRoutes register all api routers
func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("api")
	api.Use(system.JWTAuth())
	{
		v1.RegisterRouter(api)
	}
}

// RegisterOpenRoutes register routes which does not need auth
func RegisterOpenRoutes(router *gin.Engine) {
	router.POST("login", user.Login)
	router.POST("enroll", user.Enroll)
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// RegisterAppRoutes app 路由注册
func RegisterAppRoutes(router *gin.Engine) {
	app := router.Group("app")
	// 鉴权
	app.Use(system.JWTAuth())
	app.GET("hello", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello APP")
	})

}
