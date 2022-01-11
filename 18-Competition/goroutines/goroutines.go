package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello World!") // Goroutine
	write("Programming in Go!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
