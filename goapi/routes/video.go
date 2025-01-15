package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func VideoRoutes(router *gin.Engine) {
	videoGroup := router.Group("/video")
	{
		videoGroup.GET("/", getVideo)
	}
}
func getVideo(c *gin.Context) {

                // Path to your video file
                filePath := "public/1.mp4" 

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


	c.JSON(http.StatusOK, gin.H{"data": "Get all video"})
}
