package operadores

import "fmt"

func Operadores() {

	// Aritmeticos
	//==============

	var adicao int = 10 + 20
	var subtracao int = 20 - 10
	var multiplicacao int = 10 * 20
	var divisao int = 100 / 4
	var restoDivisao int = 10 / 3 // mod

	fmt.Println(adicao, subtracao, multiplicacao, divisao, restoDivisao)

	// Nao se pode fazer nenhuma operacao no Go com variaveis de tipos diferentes
	// Exemplo:
	// var numero1 int16 = 10
	// var numero2 int32 = 20
	// resultado := 10 + 20

	// Atribuicao
	//================

	var variavel1 string = "Texto"
	variavel2 := "Texto2"
	fmt.Println(variavel1, variavel2)

	// Operadores relacionais
	//==========================
	// Sao os mesmos do Java, C e C++

	// Operadores logicos
	//=============================
	// &&, ||, !

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores unarios
	//================================

	numero := 10
	numero++
	numero += 15
	numero--
	fmt.Println(numero)

	// Nao existe operadores pre-fixados em Go "--numero" por exemplo.

}
