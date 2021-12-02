package main

import "fmt"

var n int

// Can have one function per file
func init() {
	fmt.Println("Running the init function before the main function")
	n = 10
}

// Can have one function per package
func main() {
	fmt.Println("----------- Function Init -------------")
	fmt.Println(n)
}
