package routers

import (
	"dbdms/apps/user"
	"dbdms/system"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterApiRoutes(router *gin.Engine) {
	api := router.Group("api")
	api.Use(system.JWTAuth())
	api.GET("get_all_users", user.GetAllUsers)
}

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
