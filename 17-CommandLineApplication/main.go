package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting main")

	application := app.Generate()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
