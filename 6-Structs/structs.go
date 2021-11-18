package main

import (
	"fmt"
)

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	publicPlace string
	number      int8
}

func main() {
	fmt.Println("File structs")

	var u user
	u.name = "Horse"
	u.age = 30
	fmt.Println(u)

	exampleAddress := address{"Wolver street", 77}

	u2 := user{"Ox", 38, exampleAddress}
	fmt.Println(u2)

	u3 := user{age: 31}
	fmt.Println(u3)
}
