package main

import (
	"fmt"

	funcoesjson "example.com/fundamentos-adicionais/funcoes-json"
)

func main() {
	fmt.Println("Main Function")
	funcoesjson.ExemploJsonMarshal()
	funcoesjson.ExemploJsonUnmarshal()
}
