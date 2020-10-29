package main

import (
	"dbdms/urls"
	"dbdms/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func init() {
	err := utils.LoadServerConfig("config/server-config.yml")
	if err != nil {
		// helper.ErrorLogger.Errorln("Read Config file error: ", err)
		fmt.Println("read config file error")
		os.Exit(3)
	}

}

func main() {
	router := gin.New()

	// router.Use(system.Logger(helper.AccessLogger), gin.Recovery())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	ginConfig := utils.GetGinConfig()
	router.HandleMethodNotAllowed = ginConfig.HandleMethodNotAllowed
	router.Static("/page", "view")
	router.MaxMultipartMemory = ginConfig.MaxMultipartMememory
	urls.RegisterAPIRoutes(router)
	urls.RegisterOpenRoutes(router)
	// routers.RegisterAppRoutes(router)
	serverConfig := utils.GetServerConfig()
	server := &http.Server{
		Addr:           serverConfig.Addr,
		IdleTimeout:    serverConfig.IdleTimeout * time.Second,
		ReadTimeout:    serverConfig.ReadTimeout * time.Second,
		WriteTimeout:   serverConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: serverConfig.MaxHeaderBytes,
		Handler:        router,
	}
	_ = server.ListenAndServe()
}
