package main

import (
	"os"

	"dbdms/system"

	"github.com/gin-gonic/gin"
)

func init() {
	err := system.LoadServerConfig("conf/server-config.yml")
	if err != nil {
		os.Exit(3)
	}
}

func main() {
	ginConfig := system.GetGinConfig()
	gin.SetMode(ginConfig.RunMode)
	router := gin.New()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
