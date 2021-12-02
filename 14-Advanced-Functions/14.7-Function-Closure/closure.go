package main

import "fmt"

func closure() func() {
	text := "Inside the closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	fmt.Println("---------- Function Closure ----------")

	text := "Inside the main functions"
	fmt.Println(text)
	newFunction := closure()
	newFunction()
}
