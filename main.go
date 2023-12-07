package main

import (
	routes "bitbucket.org/heavenland/hl-game-loginhandler/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes.SuperRoutes(app)
	app.Run(":3030")
}
