package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	add := sum(10, 20)
	fmt.Println(add)

	result := f("Text Function f")
	fmt.Println(result)

	resultSum, _ := calculationsMathematical(10, 15)
	fmt.Println(resultSum)
}

var f = func(txt string) string {
	fmt.Println(txt)
	return txt
}

func calculationsMathematical(n1, n2 int8) (int8, int8) {
	add := n1 + n2
	subtraction := n1 - n2
	return add, subtraction
}
