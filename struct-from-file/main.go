package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Results []string
}

type Alprd struct {
	Hobbies []string
    Uuid   string `json:"uuid"`
}
func main() {
	filePath := "db.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data := Alprd{}
	err = json.Unmarshal([]byte(jsonData), &data)
	content := string(jsonData)

	// Print the contents of the file
	if err != nil {
		log.Println(err)
		return
	}
    fmt.Println(content)
	fmt.Println("Hobbies: ", data.Uuid)

}