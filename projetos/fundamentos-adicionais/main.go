package main

import (
	"fmt"

	funcoesjson "example.com/fundamentos-adicionais/funcoes-json"
	protocolohttp "example.com/fundamentos-adicionais/protocolo-http"
)

func main() {
	fmt.Println("Main Function")
	funcoesjson.ExemploJsonMarshal()
	funcoesjson.ExemploJsonUnmarshal()
	protocolohttp.ServidorHttp()
}
