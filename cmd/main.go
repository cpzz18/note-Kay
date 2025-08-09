package cmd

import (
	"github.com/gin-gonic/gin"

	"note-kay/config"
)

func main() {
	r := gin.Default()
	config.LoadEnv()
	config.ConnectDB()

	r.Run(":8080")
}
