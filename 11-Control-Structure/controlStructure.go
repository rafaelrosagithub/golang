package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------- Control Structure -------------")

	number := -5

	if number > 15 {
		fmt.Println("greater than 15")
	} else {
		fmt.Println("less than or equal to 15")
	}

	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("number is greater than 0")
	} else {
		fmt.Println("number is less than zero")
	}
}
