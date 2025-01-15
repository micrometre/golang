package main

import "github.com/gin-gonic/gin"
import "example/web-service-gin/routes" // Import the route packages

func main() {
        router := gin.Default()
        routes.UserRoutes(router)
        router.Run(":5000")
}