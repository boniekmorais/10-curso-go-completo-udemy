package funcoes

import "fmt"

func fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func ExemploFuncaoRecursiva() {

	fmt.Println("Funções Recursivas")

	// Sequencia de Fibonacci
	// 1 1 2 3 5 8 13

	posicao := uint(12)

	// fmt.Println(fibonacci(posicao))

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
