package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Alprd struct {
	Uuid string `json:"uuid"`
}

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.POST("/alprd", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := Alprd{}
		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("uuid: ", data.Uuid)
        c.JSON(http.StatusOK, gin.H{"message": "ALPRD created", "plate": data.Uuid})

	})




	router.GET("/video", func(c *gin.Context) {
		filePath := "videos/1.mp4"
		c.Header("Content-Type", "video/mp4")
		c.Header("Content-Disposition", "inline; filename=\"video.mp4\"")
		file, err := os.Open(filePath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to open file")
			return
		}
		defer file.Close()
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to stream video")
			return
		}
	})
	router.Run(":5000")
}
