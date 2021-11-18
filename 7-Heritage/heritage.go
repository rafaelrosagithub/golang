package main

import "fmt"

type people struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	people
	course  string
	college string
}

func main() {
	fmt.Println("Heritage")

	p1 := people{"João", "Peter", 21, 181}
	fmt.Println(p1)

	e1 := student{p1, "asdf", "asdf"}
	fmt.Println(e1)
	fmt.Println(e1.name)
}
