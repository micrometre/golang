package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite3 driver
)

type Alprd struct {
	Uuid    string   `json:"uuid"`
	Results []Result `json:"results"`
}

type Result struct {
	Plate      string  `json:"plate"`
}

func main() {
	// Open the SQLite database
	// Open the SQLite database
	db, err := sql.Open("sqlite3", "./alprdDb.db")
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
                        site_id TEXT
                )
        `)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}

	router.POST("/alprd", func(c *gin.Context) {
		newAlprd := struct {
			Uuid       string   `json:"uuid"`
			Confidence float64  `json:"confidence"`
			Results    []Result `json:"results"`
		}{}

		if err := c.BindJSON(&newAlprd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		stmt, err := db.Prepare("INSERT INTO users (uuid, plate) VALUES (?, ?)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(newAlprd.Uuid, newAlprd.Confidence)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(newAlprd.Results)

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	})

	router.POST("/alprd2", func(c *gin.Context) {
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
