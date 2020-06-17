package routers

import (
	"dbdms/apps/user"
	"dbdms/midware"
	v1 "dbdms/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RegisterAPIRoutes register all api routers
func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("api")
	api.Use(midware.JWTAuth())
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

// RegisterAppRoutes app 路由注册
func RegisterAppRoutes(router *gin.Engine) {
	app := router.Group("app")
	// 鉴权
	app.Use(midware.JWTAuth())
	app.GET("hello", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello APP")
	})

}
