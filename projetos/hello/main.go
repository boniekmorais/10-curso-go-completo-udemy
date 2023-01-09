package main

import (
	"fmt"
	"time"

	arraysslices "example.com/hello/pkg/arrays-slices"
	"example.com/hello/pkg/auxiliar"
	"example.com/hello/pkg/controle"
	"example.com/hello/pkg/funcoes"
	"example.com/hello/pkg/heranca"
	"example.com/hello/pkg/maps"
	"example.com/hello/pkg/metodos"
	"example.com/hello/pkg/operadores"
	"example.com/hello/pkg/ponteiros"
	"example.com/hello/pkg/structs"
	"example.com/hello/pkg/tipos"
	"example.com/hello/pkg/variaveis"
	"github.com/badoux/checkmail"
)

// Funcao iniciada com letra maiuscula pode ser exportada para outros pacotes.
// Funcao iniciada com letra minuscula somente visivel dentro do pacote.

// Função init é executada antes da função main.
// Pode ter uma função por arquivo.

func init() {
	fmt.Println("Executando a função init")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
	}
}

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

	resultado := funcoes.Somar(10, 20)
	fmt.Println("Soma:", resultado)

	funcoes.TesteFuncao()

	soma, subtracao := funcoes.CalculosMatematicos(20.0, 10.0)
	fmt.Printf("Soma.....:\t%3.2f\n", soma)
	fmt.Printf("Subtracao:\t%3.2f\n", subtracao)

	_, subtracao2 := funcoes.CalculosMatematicos(20.0, 10.0)
	fmt.Printf("Subtracao:\t%3.2f\n", subtracao2)

	operadores.Operadores()
	structs.Structs()
	heranca.Heranca()
	ponteiros.Ponteiros()
	arraysslices.ArraysSlices()
	maps.Maps()
	controle.Controle()

	somaNovo, subtracaoNovo := funcoes.CalculosMatematicos2(10, 20)

	fmt.Println(somaNovo)
	fmt.Println(subtracaoNovo)

	total := funcoes.CalculosMatematicos3(10, 20, 30)

	fmt.Println(total)

	funcoes.ExemploFuncaoAnonima()
	funcoes.ExemploFuncaoRecursiva()
	funcoes.ExemploDefer()
	funcoes.ExemploPanicRecover()
	funcoes.ExemploClosure()
	funcoes.ExemploFuncaoComPonteiros()

	metodos.ExemploMetodos()

	// Exemplo tabela
	// w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	// fmt.Fprintln(w, "a\tb\tc\td\t")
	// fmt.Fprintln(w, "aa\tbb\tcc\t")
	// fmt.Fprintln(w, "aaa\tbbb\tccc\t")
	// fmt.Fprintln(w, "aaaa\tbbbb\tcccc\tdddd\t")
	// w.Flush()

	// Prints out:
	// a    b    c    d
	// aa   bb   cc
	// aaa  bbb  ccc
	// aaaa bbbb cccc dddd

	// Exemplo table:

	// table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
	// return strings.ToUpper(fmt.Sprintf(format, vals...))
	//   }

	//   tbl := table.New("ID", "Name", "Cost ($)")

	//   for _, widget := range Widgets {
	// tbl.AddRow(widget.ID, widget.Name, widget.Cost)
	//   }

	//   tbl.Print()

	// Output:
	// ID  NAME      COST ($)
	// 1   Foobar    1.23
	// 2   Fizzbuzz  4.56
	// 3   Gizmo     78.90

}
