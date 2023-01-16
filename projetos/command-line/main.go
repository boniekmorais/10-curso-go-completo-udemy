package main

import (
	"fmt"
	"log"
	"os"

	"example.com/command-line/app"
)

func main() {
	fmt.Println("Command Line")
	aplicacao := app.Gerar()
	if error := aplicacao.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
