package routers

import (
	"dbdms/system"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterOpenRoutes(router *gin.Engine) {
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// app 路由注册
func RegisterAppRoutes(router *gin.Engine) {
	app := router.Group("app")
	// 鉴权
	app.Use(system.JWTAuth())
	app.GET("hello", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello APP")

	})

}
