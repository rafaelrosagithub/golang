package main

import "fmt"

func func1() {
	fmt.Println("Running function 1")
}

func func2() {
	fmt.Println("Running function 2")
}

func studentIsApproved(n1, n2 float32) bool {
	defer fmt.Println("Average calculated. Result will be returned")
	fmt.Println("studentIsApproved()")
	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("------------ Defer -----------")
	defer func1() // Defer
	func2()
	fmt.Println(studentIsApproved(7, 8))
}
