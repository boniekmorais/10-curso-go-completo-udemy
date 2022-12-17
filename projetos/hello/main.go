package main

import (
	"fmt"

	"example.com/hello/pkg/auxiliar"
	"example.com/hello/pkg/tipos"
	"example.com/hello/pkg/variaveis"
	"github.com/badoux/checkmail"
)

// Funcao iniciada com letra maiuscula pode ser exportada para outros pacotes.
// Funcao iniciada com letra minuscula somente visivel dentro do pacote.

func main() {
	fmt.Println("Hello there!")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("boniek.morais@gmail.com..#")

	if erro != nil {
		fmt.Println(erro)
	}

	var nome string = "John Doe"
	fmt.Println("Nome:", nome)

	endereco := "Rua Onze 20"
	fmt.Println("Endereço:", endereco)

	variaveis.ImprimirDados()

	tipos.TiposBasicos()

}
