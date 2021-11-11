package main

import (
	"errors"
	"fmt"
)

func main() {

	// Numbers
	var number int64 = 100000000000000000
	fmt.Println(number)

	var number2 uint32 = 10000 // unsygned int
	fmt.Println(number2)

	// alias
	// int32 == RUNE
	var number3 rune = 12456
	fmt.Println(number3)

	// int8 == BYTE
	var number4 byte = 123
	fmt.Println(number4)

	var realNumberFloat32 float32 = 123.45
	fmt.Println(realNumberFloat32)

	var realNumberFloat64 float64 = 12365465465.45
	fmt.Println(realNumberFloat64)

	realNumber := 123.45 // will take the architecture of the pc in the case the here it is 64 bits
	fmt.Println(realNumber)

	// Strings
	var str string = "Text1"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'R' // there is no char character in Go language
	fmt.Println(char)

	var text int32
	fmt.Println(text)

	// Boolean
	var boolean1 bool = true
	fmt.Println(boolean1)

	// Error
	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Internal error")
	fmt.Println(erro2)
}
