package funcoes

import "fmt"

func ExemploDefer() {
	fmt.Println("Exemplo defer")
	defer funcao1() // Adiar a execução da função até o último momento possível.
	funcao2()
	fmt.Println(alunoAprovado(7, 10))
}

func funcao1() {
	fmt.Println("Executando a função 1.")
}

func funcao2() {
	fmt.Println("Executando a função 2.")
}

func alunoAprovado(n1, n2 float32) bool {

	defer fmt.Println("Média calculada. Resultado será retornado") // Será executado imediatamente antes do return.

	media := (n1 + n2) / 2

	return media >= 6
}
