package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna a resposta em JSON para a requisição.
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}

}

// Retorna o erro no formato JSON.
func Erro(w http.ResponseWriter, statusCode int, erro error) {

	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
