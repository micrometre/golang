package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
   Name    string
   Age     int64
   Hobbies []string
}

func main() {
   filePath := "db.json" 
   jsonData, err := os.ReadFile(filePath) 
   jsonStr := "{\"name\": \"Prashant\", \"age\": 27, \"hobbies\":[\"sports\",\"music\"]}"
   data := Person{}
   err = json.Unmarshal([]byte(jsonStr), &data)
       content := string(jsonData)

        // Print the contents of the file
        fmt.Println(content)
   if err != nil {
      log.Println(err)
      return
   }

   fmt.Println(data.Name)
}