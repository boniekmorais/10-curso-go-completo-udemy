package tipos

import (
	"errors"
	"fmt"
)

func TiposBasicos() {

	// int8, int16, int32, int64
	// int => usa arquitetura do computador
	// Arquitetura 32 bits usa int32
	// Arquitetura 64 bits usa int64

	// uint -> unsigned int

	var numero int16 = 100
	fmt.Println("Numero:", numero)

	// Alias
	// int32 = rune
	// uint8 = byte

	var numero3 rune = 200
	fmt.Println("Numero alias:", numero3)

	var numero4 byte = 200
	fmt.Println("Numero alias2:", numero4)

	// float32, float64

	var numero5 float64 = 200.9
	fmt.Println("Numero float:", numero5)

	var texto string = "Este é um texto"
	fmt.Println("String:", texto)

	// Imprime o valor numerico de B da tabela ASCII

	char := 'B'
	fmt.Println("Char:", char)

	// Valor zero

	var texto2 string
	fmt.Println("String em branco", texto2)

	// Valores iniciais
	// String: vazio
	// int e float: 0
	// bool: false
	// error: nil

	// boolean

	var test bool
	fmt.Println("Boolean:", test)

	var erro error
	fmt.Println("Error:", erro)

	var erro2 error = errors.New("erro interno")
	fmt.Println("Error:", erro2)

}
