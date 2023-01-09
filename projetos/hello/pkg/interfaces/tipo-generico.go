package interfaces

import "fmt"

func ExemploTipoGenerico() {
	fmt.Println("Tipo genérico")
	generica("Texto")
	generica(109)
	generica(true)
}

func generica(i interface{}) {
	fmt.Println(i)
}
