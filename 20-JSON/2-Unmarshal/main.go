package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	dogInJSON := `{"name":"Rex","breed":"Dálmata","age":5}`

	var d dog

	if erro := json.Unmarshal([]byte(dogInJSON), &d); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(d)

	dog2InJSON := `{"name":  "Toby", "breed": "Poodle"}`

	d2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(dog2InJSON), &d2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(d2)
}
