package channelsbuffer

import "fmt"

func ChannelsBuffer() {
	fmt.Println("Channels with buffer")

	canal := make(chan string, 2) // Buffer. Capacidade para o canal. Só bloqueia ao atingir capacidade máxima.
	canal <- "Hello"
	canal <- "Hello again"
	// canal <- "Hello deadlock" // Se utilizar acima do buffer, ocorre deadlock.

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
