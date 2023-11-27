package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Retorna um router com as portas configuradas.
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
