package main

import "fmt"

func main() {
	fmt.Println("--------- Anonymous functions ---------")

	func() {
		fmt.Println("Hello World!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Hello World!!!")

	returnFunction := func(text string) string {
		return fmt.Sprintf("Received: %s", text)

	}("Passing parameter")

	fmt.Println(returnFunction)

}
