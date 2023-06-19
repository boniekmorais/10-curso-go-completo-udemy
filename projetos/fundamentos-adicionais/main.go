package main

import (
	"fmt"

	protocolohttp "example.com/fundamentos-adicionais/protocolo-http"
)

func main() {
	fmt.Println("Main Function")
	// funcoesjson.ExemploJsonMarshal()
	// funcoesjson.ExemploJsonUnmarshal()
	protocolohttp.ServidorHttp()
}
