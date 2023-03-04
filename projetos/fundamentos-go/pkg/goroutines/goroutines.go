package goroutines

import (
	"fmt"
	"time"
)

func GoRoutines() {
	// Concorrencia != Paralelismo.
	fmt.Println("Go Routines")
	go escrever("Hello World!") // Go Routine
	escrever(("Programando em Go!"))
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
