package routes

import (
	"fmt"
	"net/http"
	"time"

	controllers "bitbucket.org/heavenland/hl-game-loginhandler/controllers"
	"github.com/gin-gonic/gin"
)

func default_logger(param gin.LogFormatterParams) string {

	// your custom format
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func SuperRoutes(g *gin.Engine) {
	g.Use(gin.LoggerWithFormatter(default_logger))
	baseEndpoint(g)
	routesPaths(g)
}

func routesPaths(g *gin.Engine) {
	routePath := g.Group("/api")
	{
		routePath.GET("/test", controllers.TestController())
		routePath.GET("/unmarshall", controllers.TestMarhsallUnMarshall())
	}
	// Middleware if required
}

func baseEndpoint(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Response": "Websocket API Login",
			"Message":  "Base",
		})
	})
}

// Note this Handler function can be returned as well
// @Petr this is an example for how to make actual function, it is similar with Express JS and Middleware can be a function with c.Next() at the end of the return
// This is just a example base,
// All the time request and responses are handled by gin.Context
// While endpoint and other server related stuff handled by gin.Engine
func returnableEndpointFunction() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Respond": "Base function",
			"Message": "Base Function",
		})
	}
}
