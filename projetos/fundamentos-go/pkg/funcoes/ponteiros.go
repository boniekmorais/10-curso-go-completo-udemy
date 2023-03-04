package funcoes

import "fmt"

func ExemploFuncaoComPonteiros() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)
	inverterSinalPonteiro(&numero)
	fmt.Println(numero)
}

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}
