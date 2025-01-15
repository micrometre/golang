package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite3 driver
)



type Alprd struct {
    Uuid       string   `json:"uuid"`
    Results    []Result `json:"results"`
}

// Define the Result struct for the results within the Alprd struct
type Result struct {
    Plate string `json:"plate"`
}




func main() {
	db, err := sql.Open("sqlite3", "./mydatabase.db")
    if err != nil {
        log.Fatal("failed to open database: ", err)
    }
    defer db.Close()


	// Create a Gin router
	router := gin.Default()
	// Create a sample table (if not exists)
	_, err = db.Exec(`
                CREATE TABLE IF NOT EXISTS users (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        uuid TEXT,
                        plate TEXT
                )
        `)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}


    router.POST("/alprd", func(c *gin.Context) {
        var alprd Alprd

        if err := c.BindJSON(&alprd); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
            return
        }

        stmt, err := db.Prepare("INSERT INTO users (uuid, plate) VALUES (?, ?)")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer stmt.Close()

        var plate interface{}
        if alprd.Results == nil || len(alprd.Results) <= 0 {
            plate = nil // Set plate to nil if no results are present
        } else {
            plate = alprd.Results[0].Plate // Extract the plate from the first result
        }

        _, err = stmt.Exec(alprd.Uuid, plate)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

		fmt.Println(plate)
        c.JSON(http.StatusCreated, gin.H{"message": "Data inserted successfully"})
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
