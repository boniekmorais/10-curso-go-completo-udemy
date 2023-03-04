package funcoes

import "fmt"

func Somar(a int, b int) int {
	return a + b
}

func TesteFuncao() {

	var f = func(arg string) string {
		fmt.Println("Variável função:", arg)
		return arg
	}

	f("Teste")

	resultado := f("Resultado")
	fmt.Println(resultado)
}

func CalculosMatematicos(a float64, b float64) (float64, float64) {
	soma := a + b
	subtracao := a - b
	return soma, subtracao
}
