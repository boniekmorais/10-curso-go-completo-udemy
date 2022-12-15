package variaveis

import (
	"fmt"
)

func ImprimirDados() {
	var nome string = "John Doe"
	fmt.Println("Nome:", nome)

	endereco := "Rua Onze 20"
	fmt.Println("Endereço:", endereco)

	// Declarando varias variaveis de uma unica vez:

	var (
		idade  int
		cidade string
		estado string
	)

	idade = 39
	cidade = "Divinopolis"
	estado = "Minas Gerais"

	fmt.Println("Idade:", idade)
	fmt.Println("Cidade:", cidade)
	fmt.Println("Estado:", estado)

	empresa, matriz := "Sensedia", "Campinas"

	fmt.Println("Empresa:", empresa)
	fmt.Println("Matriz:", matriz)

	// Constantes

	const TIPO_CARGO string = "Software Developer"
	fmt.Println("Cargo:", TIPO_CARGO)

	// Invertendo valor de variaveis sem uso de variavel auxiliar

	empresa, matriz = matriz, empresa

	fmt.Println("Empresa:", empresa)
	fmt.Println("Matriz:", matriz)

}
