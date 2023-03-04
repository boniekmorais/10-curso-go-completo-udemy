package channels

import (
	"fmt"
	"time"
)

// Channels: Podem enviar e receber dados entre Go Routines.

func Channels() {
	fmt.Println("Channels")

	canal := make(chan string)

	go escrever("Hello World!", canal)

	fmt.Println("Depois da execução da função escrever.")

	for {
		mensagem, aberto := <-canal // Canal aguardando receber o valor.

		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Canal recebendo o valor.
		time.Sleep(time.Second)
	}

	close(canal)
}
