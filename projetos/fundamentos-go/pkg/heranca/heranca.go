package heranca

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

// Pega todos os campos da struct 'pessoa' e usa dentro da struct 'estudante'.

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func Heranca() {

	p1 := pessoa{
		nome:  "Boniek",
		idade: 39,
	}

	aluno := estudante{
		pessoa:    p1,
		curso:     "Computação",
		faculdade: "PUC",
	}

	fmt.Println(aluno.nome)
	fmt.Println(aluno.idade)
	fmt.Println(aluno)

	p2 := pessoa{"Boniek", 39}
	a2 := estudante{p2, "Computação", "PUC"}

	fmt.Println(a2.nome)
	fmt.Println(a2.idade)
	fmt.Println(a2)

}
