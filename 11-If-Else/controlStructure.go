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
	} else if number < -10 {
		fmt.Println("number is less than -10")
	} else {
		fmt.Println("between 0 and -10")
	}
}
