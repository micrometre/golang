package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)



type Alprd struct {
    Uuid   string `json:"uuid"`
    Results []Result `json:"results"`
}


type Result struct {
    Plate   string `json:"plate"`
    Confidence   float64 `json:"confidence"`
}


func main() {
	filePath := "db.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data := Alprd{}
	err = json.Unmarshal([]byte(jsonData), &data)
	//content := string(jsonData)

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("uuid: ", data.Uuid)
	fmt.Println("Plate: ", data.Results)
}