package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("---------- Generic Types ------------")
	generic("String")
	generic(7)
	generic(true)

	fmt.Println(1, 2, "String", false, true, float64(1234567))

	// It's not a good practice
	maps := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(maps)
}
