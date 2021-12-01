package routes

import (
	"go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.GET("/user", controllers.User)
		v1.POST("/logout", controllers.Logout)
	}
	return r
}
