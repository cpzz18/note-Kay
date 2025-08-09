package cmd

import (
	"github.com/gin-gonic/gin"
	"honnef.co/go/tools/config"

	"note-kay/routes"
)

func main() {
	r := gin.Default()
	config.LoadEnv()
	config.Connect
	routes.Setup
}
