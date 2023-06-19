package protocolohttp

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func ServidorHttp() {

	templates = template.Must(template.ParseGlob("*.html"))

	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Olá Mundo!"))
	// })

	// http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Olá usuários!"))
	// })

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", nil)
	})

	log.Fatal(http.ListenAndServe(":5001", nil))
}
