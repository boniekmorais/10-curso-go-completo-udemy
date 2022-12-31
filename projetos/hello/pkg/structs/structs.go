package structs

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func Structs() {

	fmt.Println("Structs")

	var usuario1 usuario
	var usuario2 usuario

	usuario1.nome = "Boniek Morais"
	usuario1.idade = 39

	usuario2.nome = "John Doe"
	usuario2.idade = 45

	fmt.Println(usuario1)
	fmt.Println(usuario2)

	enderecoExemplo := endereco{"Rua Onze", 20}

	usuario3 := usuario{"Barbara", 23, enderecoExemplo}
	fmt.Println(usuario3)

	usuario4 := usuario{nome: "Carlos"}
	fmt.Println(usuario4)

}
