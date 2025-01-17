package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite3 driver
)

type Alprd struct {
	Uuid    string   `json:"uuid"`
	Results []Result `json:"results"`
}

// Define the Result struct for the results within the Alprd struct
type Result struct {
	Plate string `json:"plate"`
}

func main() {
	db, err := sql.Open("sqlite3", "data/sqlite.db")
	if err != nil {
		log.Fatal("failed to open database: ", err)
	}
	defer db.Close()

	// Create a Gin router
	router := gin.Default()
	router.Static("/public", "./public")
	// Create a sample table (if not exists)
	_, err = db.Exec(`
                CREATE TABLE IF NOT EXISTS alprd (
                        plate TEXT PRIMARY KEY,
                        uuid TEXT
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
		stmt, err := db.Prepare("INSERT INTO alprd (plate, uuid) VALUES (?, ?)")
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
		var imageUrl string = "http://127.0.0.1:5000/public/images/" + alprd.Uuid + ".jpg"

		_, err = stmt.Exec(plate, imageUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(imageUrl)
		c.JSON(http.StatusCreated, gin.H{"message": "Data inserted successfully"})
	})

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
		cmd := exec.Command("alpr", "-c gb", filename) // Replace with the actual command

		// Capture output
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "stderr": stderr.String()})
			return
		}

		output := stdout.String()
		fmt.Println(output)
		c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
	})

	router.GET("/video", func(c *gin.Context) {
		filePath := "public/videos/1.mp4"
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
