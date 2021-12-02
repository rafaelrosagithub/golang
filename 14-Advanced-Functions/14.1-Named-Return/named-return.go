package main

import "fmt"

func mathematicalCalculations(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	fmt.Println("---------- Named return ------------")

	sum, subtraction := mathematicalCalculations(10, 20)
	fmt.Println(sum, subtraction)
}
