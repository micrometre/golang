package main

import (
	"example/web-service-gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Import the route packages

func main() {
        router := gin.Default()
        routes.UserRoutes(router)
        routes.VideoRoutes(router)
        router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello woooooooo        rld")
	})
        router.Run(":5000")
}