package main

import (
	helper "dbdms/helpers"
	"dbdms/routers"
	system "dbdms/system"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	err := system.LoadServerConfig("conf/server-config.yml")
	if err != nil {
		helper.ErrorLogger.Errorln("Read Config file error: ", err)
		os.Exit(3)
	}
}

func main() {
	ginConfig := system.GetGinConfig()
	gin.SetMode(ginConfig.RunMode)
	router := gin.New()
	router.Use(system.Logger(helper.AccessLogger), gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	router.HandleMethodNotAllowed = ginConfig.HandleMethodNotAllowed
	router.Static("/page", "view")
	router.MaxMultipartMemory = ginConfig.MaxMultipartMememory
	routers.RegisterApiRoutes(router)
	serverConfig := system.GetServerConfig()
	server := &http.Server{
		Addr:           serverConfig.Addr,
		IdleTimeout:    serverConfig.IdleTimeout * time.Second,
		ReadTimeout:    serverConfig.ReadTimeout * time.Second,
		WriteTimeout:   serverConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: serverConfig.MaxHeaderBytes,
		Handler:        router,
	}
	server.ListenAndServe()
}
