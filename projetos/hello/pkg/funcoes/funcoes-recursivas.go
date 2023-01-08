package funcoes

import "fmt"

func ExemploFuncaoRecursiva() {

	fmt.Println("Funções Recursivas")

	// Sequencia de Fibonacci
	// 1 1 2 3 5 8 13

	posicao := uint(10)

	fmt.Println(fibonacci(posicao))

}

func fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
