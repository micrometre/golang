package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Hobbies []string
}

func main() {
	filePath := "db.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr := "{\"name\": \"Prashant\", \"age\": 27, \"hobbies\":[\"sports\",\"music\"]}"
	data := Person{}
	err = json.Unmarshal([]byte(jsonStr), &data)
	content := string(jsonData)

	// Print the contents of the file
	if err != nil {
		log.Println(err)
		return
	}
    fmt.Println((content))
}