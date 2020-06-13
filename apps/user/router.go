package user

import "github.com/gin-gonic/gin"

// RegisterRouter use router register
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("", GetAllUsers)
	router.GET("/:id", GetUserProfile)
	router.DELETE("/:id", DeleteUser)
	router.PUT("/:id", UpdateUserProfile)
}
