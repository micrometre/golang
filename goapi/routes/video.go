// routes/products.go

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideoRoutes(router *gin.Engine) {
	videoGroup := router.Group("/video")
	{
		videoGroup.GET("/", getvideo)
	}
}

func getvideo(c *gin.Context) {
	// ... your logic to get all products
	c.JSON(http.StatusOK, gin.H{"data": "Get all video"})
}
