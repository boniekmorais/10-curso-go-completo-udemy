package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

// Waitgroup é usado para sincronizar duas ou mais Go Routines.

func Waitgroup() {

	fmt.Println("Waitgroup")

	var waitgroup sync.WaitGroup

	waitgroup.Add(2)

	go func() {
		escrever("Hello World!")
		waitgroup.Done()
	}()

	go func() {
		escrever(("Programando em Go!"))
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
