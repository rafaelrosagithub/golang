package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("Hello world!", channel)

	fmt.Println("After write")

	// for {
	// 	message, open := <-channel
	// 	if !open {
	// 		fmt.Println("Closed channel")
	// 		break
	// 	}
	// 	fmt.Println("After message := <-channel")
	// 	fmt.Println(message)
	// }

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("End of program!")
}

func write(text string, channel chan string) {
	fmt.Println("func write()")
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
