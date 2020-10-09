package jwtauth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AUTH jwt user auth midware
func AUTH() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("JWT AUTH midware")
		path := context.Request.URL.Path
		if strings.Contains(path, "swagger") {
			return
		}

		token := context.Request.Header.Get("ACCESS_TOKEN")
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": "permission denied, Request has no token",
			})
			context.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ResolveToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": err.Error(),
			})
			context.Abort()
			return
		}
		context.Set("claims", claims)
		context.Next()
	}
}
