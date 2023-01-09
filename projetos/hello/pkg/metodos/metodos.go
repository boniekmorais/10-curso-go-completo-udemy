package metodos

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Println("Salvando usuario: ", u.nome)
}

func (u *usuario) incrementarIdade() {
	u.idade++
}

func ExemploMetodos() {
	fmt.Println("Métodos")

	funcionario := usuario{
		nome:  "Boniek",
		idade: 39,
	}

	funcionario2 := usuario{"John Doe", 33}

	funcionario.salvar()
	funcionario2.salvar()

	funcionario.incrementarIdade()
	fmt.Println("Idade: ", funcionario.idade)
}
