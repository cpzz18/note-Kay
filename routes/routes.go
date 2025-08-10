package routes

import (
	"github.com/gin-gonic/gin"

	"note-kay/controllers"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

	}
}
