package routes

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import the sqlite3 driver
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

func AlprdRoutes(router *gin.Engine) {

	alprdGroup := router.Group("/alprd")
	{
		alprdGroup.POST("/", getAlprd)
	}
}
func getAlprd(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./sqlitedatabase.db")
	if err != nil {
		log.Fatal("failed to open database: ", err)
	}
	defer db.Close()
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
	var alprd Alprd

	if err := c.BindJSON(&alprd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	fmt.Println(c)
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

}
