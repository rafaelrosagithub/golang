package main

import "fmt"

func reverseSign(number int) int {
	return number * -1
}

func reverseSignWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("----------- Pointers -----------")

	number := 20
	reverseSign := reverseSign(number)
	fmt.Println(reverseSign)

	newNumber := 40
	fmt.Println(newNumber)
	reverseSignWithPointer(&newNumber)
	fmt.Println(newNumber)
}
