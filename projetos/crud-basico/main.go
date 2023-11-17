package main

import (
	"fmt"
	"log"
	"net/http"

	userService "example.com/crud-basico/service"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", userService.CreateUser).Methods(http.MethodPost)

	fmt.Println("Starting web server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
