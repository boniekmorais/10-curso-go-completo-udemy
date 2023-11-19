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
	router.HandleFunc("/users", userService.ListUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userService.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userService.UpdateUserById).Methods(http.MethodPut)

	fmt.Println("Starting web server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
