package funcoes

// Número variável de argumentos

func CalculosMatematicos3(numeros ...int) int {

	soma := 0

	for _, n := range numeros {
		soma += n
	}

	return soma

}
