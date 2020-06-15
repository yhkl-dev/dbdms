package system

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth Mid ware
func Permission() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if strings.Contains(path, "swagger") {
			return
		}
		fmt.Println(path)
		method := context.Request.Method
		fmt.Println(method)
	}

}
