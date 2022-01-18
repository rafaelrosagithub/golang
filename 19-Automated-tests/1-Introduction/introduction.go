package main

import (
	"fmt"
	"test-introduction/adresses"
)

func main() {
	addressType := adresses.AddressType("Road")
	fmt.Println(addressType)
}
