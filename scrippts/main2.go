package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)



type ALPRD struct {
	plate  string `form:"plate" binding:"required"`
}
func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Video streaming platform!")
	})

	router.GET("/video", func(c *gin.Context) {
		// Path to your video file
		filePath := "videos/1.mp4"
		// Set response headers for video streaming
		c.Header("Content-Type", "video/mp4")
		c.Header("Content-Disposition", "inline; filename=\"video.mp4\"")
		// Open the video file
		file, err := os.Open(filePath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to open file")
			return
		}
		defer file.Close()
		// Stream the video data
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to stream video")
			return
		}
	})


	router.POST("/alprd", func(c *gin.Context) {
		var alprd ALPRD
		if err := c.ShouldBind(&alprd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		println(c.Request)
		c.JSON(http.StatusOK, gin.H{"message": "ALPRD created", "plate": alprd.plate})
	})


	router.Run(":5000")
}
