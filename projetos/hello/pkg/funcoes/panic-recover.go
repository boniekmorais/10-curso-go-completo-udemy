package funcoes

import "fmt"

func ExemploPanicRecover() {

	fmt.Println(alunoAprovado2(6, 6))
	fmt.Println("Pós execução")

}

func alunoAprovado2(n1, n2 float32) bool {

	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!") // Interrompe a execução do programa. Antes chama todas as funções que tem defer.

}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}
