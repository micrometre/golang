package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite3 driver
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Alprd struct {
	Uuid    string   `json:"uuid"`
	Results []Result `json:"results"`
	Plate string `json:"plate"`

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
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*")

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

	
	router.GET("/", func(c *gin.Context) {
  c.HTML(
      http.StatusOK,
      "index.html",
      gin.H{
          "title": "Home Page",
      },
  )

})

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
		// Get the uploaded file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting uploaded file: " + err.Error()})
			return
		}
		// Validate file extension (optional)
		if !isAllowedExtension(file.Filename) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension. Only specific extensions allowed"})
			return
		}
		// Generate a unique filename
		filename := filepath.Base(file.Filename)
		newFilename := filename
		// Save the uploaded file
		if err := c.SaveUploadedFile(file, "public/"+newFilename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving uploaded file: " + err.Error()})
			return
		}
		// Execute ffmpeg command
		cmd := exec.Command("ffmpeg", "-i", "public/"+newFilename, "-listen", "1", "-f", "mp4", "-movflags", "frag_keyframe+empty_moov", "http://127.0.0.1:5001")
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "stderr": stderr.String()})
			return
		}
		output := stdout.String()
		fmt.Println(newFilename, output)
		c.JSON(http.StatusOK, gin.H{"message": "File uploaded and processed successfully"})
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



    // New GET route to retrieve all ALPR data
    router.GET("/alprd", func(c *gin.Context) {
        rows, err := db.Query("SELECT * FROM alprd")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        var alprds []Alprd
        for rows.Next() {
            var alprd Alprd
            err := rows.Scan(&alprd.Plate, &alprd.Uuid)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            alprds = append(alprds, alprd)
        }

        c.JSON(http.StatusOK, alprds)
    })

	router.Run(":5000")
}

// Function to check for allowed file extensions (optional)
func isAllowedExtension(filename string) bool {
	allowedExtensions := map[string]bool{
		".mp4": true,
		".avi": true,
		// Add more allowed extensions here
	}
	ext := filepath.Ext(filename)
	return allowedExtensions[ext]
}
