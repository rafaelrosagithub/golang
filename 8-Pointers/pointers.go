package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// Poiter is a memory reference
	var var3 int
	var pointer *int

	var3 = 100
	pointer = &var3

	fmt.Println(var3, pointer) // pointer para obter o endereço de memória

	var3 = 150
	fmt.Println(var3, *pointer) // *pointer pra obter o valor
}
