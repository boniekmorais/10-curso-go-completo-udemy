package funcoes

import "fmt"

func ExemploFuncaoAnonima() {

	func() {
		fmt.Println("Hello there!")
	}()

	resultado := func(texto string) string {
		return fmt.Sprintf("Texto recebido: %s", texto)
	}("John Doe")

	fmt.Println(resultado)
}
