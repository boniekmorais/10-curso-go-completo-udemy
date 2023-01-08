package controle

import (
	"fmt"
	"time"
)

func Controle() {
	fmt.Println("Estruturas de Controle")

	// IF
	//--------------------------------------------------------------

	numero := 10

	if numero > 15 {
		fmt.Println("Número maior que 15")
	} else {
		fmt.Println("Número menor ou igual a 15")
	}

	// if init. Limita a variável no escopo do if.

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que -10")
	}

	// SWITCH
	// -------------------------------------------------------------

	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana2(6))
	fmt.Println(diaDaSemana3(3))

	// LOOP
	// -------------------------------------------------------------

	// Apenas for. Não tem do/while ou while.

	i := 0

	for i < 3 {
		time.Sleep(time.Second)
		i++
	}

	fmt.Println(i)

	for j := 0; j < 3; j++ {
		fmt.Println("Value:", j)
		time.Sleep(time.Second)
	}

	for j := 0; j < 100; j += 2 {
		fmt.Println("Value:", j)
		time.Sleep(time.Microsecond)
	}

	names := [3]string{"Carlos", "John", "Paul"}

	for index, name := range names {
		fmt.Println(name, index)
	}

	// Imprime cada letra da palavra com código ASCII.

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	// Imprime cada letra da palavra.

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"name":  "John Doe",
		"State": "Ohio",
	}

	for key, value := range usuario {
		fmt.Println(key, value)
	}

	// Loop infinito.

	// for {
	// 	fmt.Println("Executando infinitamente.")
	// 	time.Sleep(time.Second)
	// }

}

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Error"
	}
}

func diaDaSemana2(numero int) string {

	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Error"
	}
}

func diaDaSemana3(numero int) string {

	// Não existe o comando break no switch.

	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough // Força cair na próxima condição sem avaliar. Se receber 1 será retornado "Segunda-Feira".
	case numero == 2:
		dia = "Segunda-Feira"
	case numero == 3:
		dia = "Terça-Feira"
	case numero == 4:
		dia = "Quarta-Feira"
	case numero == 5:
		dia = "Quinta-Feira"
	case numero == 6:
		dia = "Sexta-Feira"
	case numero == 7:
		dia = "Sábado"
	default:
		dia = "Error"
	}

	return dia
}
