package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Arrays
	var array [5]string
	fmt.Println(array)

	array[1] = "Position 2"
	fmt.Print(array[1])

	array2 := [3]string{"Position 1", "Position 2", "Position 3"}
	fmt.Println(array2)

	arry3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arry3)

	//Slices
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println((slice))

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 19)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Altered position"
	fmt.Println(slice2)
}
