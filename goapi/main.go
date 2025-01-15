package main

import (
	"example/web-service-gin/routes"

	"github.com/gin-gonic/gin"
)

// Import the route packages

func main() {
	router := gin.Default()
	router.Static("/", "./public")

	routes.UserRoutes(router)
	router.Run(":5000")
}
