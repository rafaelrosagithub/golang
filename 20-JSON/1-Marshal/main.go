package main

import (
	"bytes"
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
	d := dog{"Rex", "Dálmata", 5}
	fmt.Println(d)

	dogInJSON, erro := json.Marshal(d)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dogInJSON)
	fmt.Println(bytes.NewBuffer(dogInJSON))

	d2 := map[string]string{
		"name":  "Toby",
		"breed": "Poodle",
	}

	dog2InJSON, erro := json.Marshal(d2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dog2InJSON)
	fmt.Println(bytes.NewBuffer(dog2InJSON))

}
