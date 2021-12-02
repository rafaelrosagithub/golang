package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println("---------- Variatic functions ------------")
	totalSum := sum(1, 2, 3, 10, 21)
	fmt.Println(totalSum)

	write("Hello World!: ", 1, 4, 7)
}
