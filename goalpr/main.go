package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
type ALPRD struct {
    Plate string `json:"plate"` // Field name should match the JSON key
}



func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")


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

    // POST route handler
    router.POST("/alprd", func(c *gin.Context) {
        var alprd ALPRD

        // Read the request body
        body, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
            return
        }

        // Unmarshal the request body into the ALPRD struct
        err = json.Unmarshal(body, &alprd)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
            return
        }

        // Capture the plate number
        plate := alprd.Plate

        // Process the plate number (optional)
        // You can perform additional validation, filtering, or logic here

        // Send a successful response with the plate number
        c.JSON(http.StatusOK, gin.H{"message": "ALPRD created", "plate": plate})
    })



	router.Run(":5000")
}
