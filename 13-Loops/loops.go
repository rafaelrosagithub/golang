package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---------- Loops ----------")

	i := 0

	for i < 2 {
		i++
		fmt.Println("Incrementing i: ", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 2; j++ {
		fmt.Println("Incrementing j: ", j)
		time.Sleep(time.Second)
	}

	for n := 0; n < 10; n += 2 {
		fmt.Println("Incrementing j: ", n)
		time.Sleep(time.Second)
	}

	names := [3]string{"Dog", "Cat", "Ox"}

	for index, name := range names {
		fmt.Println("Idex: ", index, "Name: ", name)
	}

	for _, name := range names {
		fmt.Println("Name: ", name)
	}

	for index, letter := range "Word" {
		fmt.Println(index, letter, string(letter))
	}

	user := map[string]string{
		"name":    "Jhoanas",
		"surname": "Kepler",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	// Running endlessly
	// for {
	// 	fmt.Println("Infinite Loop")
	// 	time.Sleep(time.Second)
	// }
}
