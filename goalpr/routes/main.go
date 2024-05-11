package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to our video streaming platform!")
	})

	router.GET("/stream/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		file, err := os.Open("videos/" + filename)
		if err != nil {
			c.String(http.StatusNotFound, "Video not found.")
			return
		}
		defer file.Close()

		c.Header("Content-Type", "video/mp4")
		buffer := make([]byte, 64*1024) // 64KB buffer size
		io.CopyBuffer(c.Writer, file, buffer)
	})

	router.POST("/alprd", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		println(string(body))
	})

	router.Run(":5000")
})
		v1.POST("/	router.Run(":8080")
}ackage app
