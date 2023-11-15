package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Users"))
}

func main() {

	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })

	// http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Users"))
	// })

	// Usa a funcao declarada acima
	http.HandleFunc("/home", home)

	// Usa a funcao declarada acima
	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5001", nil))
}
