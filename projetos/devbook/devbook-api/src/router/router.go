package router

import "github.com/gorilla/mux"

// Retorna um router com as portas configuradas.
func Gerar() *mux.Router {
	return mux.NewRouter()
}
