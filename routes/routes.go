package routes

import (
	"github.com/gin-gonic/gin"

	"note-kay/controllers"
	"note-kay/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Notes
	notes := r.Group("/notes")
	notes.Use(middleware.JwtAuth())
	{
		notes.POST("/", controllers.CreateNotes)
		notes.GET("/", controllers.GetNotes)
		notes.GET("/:id", controllers.GetNote)
		notes.PUT("/:id", controllers.UpdateNote)
		notes.DELETE("/:id", controllers.DeleteNote)
	}

	// Folders
	folders := r.Group("/folders")
	folders.Use(middleware.JwtAuth())
	{
		folders.POST("/", controllers.CreateFolder)
		folders.GET("/", controllers.GetFolders)
		folders.GET("/:id", controllers.GetFolders)
		folders.PUT("/:id", controllers.UpdateFolder)
		folders.DELETE("/:id", controllers.DeleteFolder)
	}

	// Tags
	tags := r.Group("/tags")
	tags.Use(middleware.JwtAuth())
	{
		tags.POST("/", controllers.CreateTag)
		tags.GET("/", controllers.GetTags)
		tags.GET("/:id", controllers.UpdateTag)
		tags.DELETE("/:id", controllers.DeleteTag)
	}
}
