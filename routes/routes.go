package routes

import (
	"github.com/gin-gonic/gin"

	"note-kay/controllers"
	"note-kay/middleware"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/")
	api.Use(middleware.JwtAuth())
	{
		// Notes
		notes := api.Group("/notes")
		{
			notes.POST("/", controllers.CreateNotes)
			notes.GET("/", controllers.GetNotes)
			notes.GET("/:id", controllers.GetNote)
			notes.PUT("/:id", controllers.UpdateNote)
			notes.DELETE("/:id", controllers.DeleteNote)
		}

		// Folders
		folders := api.Group("/folders")
		{
			folders.POST("/", controllers.CreateFolder)
			folders.GET("/", controllers.GetFolders)
			folders.GET("/:id", controllers.GetFolders)
			folders.PUT("/:id", controllers.UpdateFolder)
			folders.DELETE("/:id", controllers.DeleteFolder)
		}

		// Tags
		tags := api.Group("/tags")
		{
			tags.POST("/", controllers.CreateTag)
			tags.GET("/", controllers.GetTags)
			tags.PUT("/:id", controllers.UpdateTag)
			tags.DELETE("/:id", controllers.DeleteTag)
		}
	}
}
