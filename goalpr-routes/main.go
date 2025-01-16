package main

import (
	"example/web-service-gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Import the route packages

func main() {


	router := gin.Default()
	router.Static("/public", "./public")
	// Create a sample table (if not exists)


	routes.UserRoutes(router)
	routes.VideoRoutes(router)
	routes.AlprdRoutes(router)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	router.Run(":5000")
}
