package main

import "fmt"

func main() {

	// Arithmetic
	sum := 1 + 2
	subtraction := 1 - 2
	division := 10 / 4
	multiplication := 10 * 5
	remainingDebt := 10 % 2 // Resto da divisão

	fmt.Println(sum, subtraction, division, multiplication, remainingDebt)

	var num1 int16 = 10
	var num2 int16 = 30
	sum2 := num1 + num2
	fmt.Println(sum2)

	// Assignment
	fmt.Println("-----------------------------")
	var var1 string = "String"
	var2 := "String2"
	fmt.Println(var1, var2)

	// Relational operators
	fmt.Println("-----------------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Logical operators
	fmt.Println("-----------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Unary operators
	fmt.Println("-----------------------------")
	number := 10
	number++
	number += 15
	fmt.Println(number)

	number--
	number -= 20

	number *= 3
	number /= 10
	number %= 3
	fmt.Println(number)
}
