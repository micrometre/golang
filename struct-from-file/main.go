package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)
type Results struct {
    Results []Result `json:"users"`
}

type Result struct {
    Plate   string `json:"plate"`

}


type Alprd struct {
    Uuid   string `json:"uuid"`
    Plate   string `json:"plate"`
    Results []Result `json:"results"`
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
	fmt.Println("uuid: ", data.Results)
}