package main // go-json

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonBlob := []byte(`[
	{"name": "John", "info":{"age": 25} },
	{"name": "Michael", "info":{"age": 30} }
]`)
	type Employee struct {
		Name string
		Data struct {
			Age int
		} `json:"info"`
	}
	var employees []Employee
	err := json.Unmarshal(jsonBlob, &employees)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Parsed: %+v\n", employees)

	newJson, err := json.Marshal(employees)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Back to json:", string(newJson))
}
