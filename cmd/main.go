package main

import (
	"github.com/gin-gonic/gin"

	"note-kay/config"
	"note-kay/routes"
)

func main() {
	r := gin.Default()

	config.LoadEnv()

	config.ConnectDB()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
