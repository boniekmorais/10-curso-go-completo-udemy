package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("DevBook")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
