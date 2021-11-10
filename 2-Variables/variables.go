package main

import "fmt"

func main() {
	var variable1 string = "Variable 1"
	variable2 := "Variable 2"
	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		var3 string = "var 3"
		var4 string = "var 4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "var 5", "var 6"
	fmt.Println(var5, var6)

	const constante1 string = "Const1"
	fmt.Println(constante1)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}
