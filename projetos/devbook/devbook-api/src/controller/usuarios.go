package controller

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Criar usuário.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	requestBody, error := io.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var usuario model.Usuario

	if error = json.Unmarshal(requestBody, &usuario); error != nil {
		log.Fatal(error)
	}

	db, error := database.Conectar()

	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	usuarioRepository := repository.NovoRepositorioUsuarios(db)
	usuarioID, erro := usuarioRepository.Criar(usuario)

	if erro != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))

}

// Listar todos os usuários.
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar usuários"))
}

// Listar usuário por ID.
func ListarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar usuário por ID"))
}

// Atualizar usuário.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuário"))
}

// Remover usuário.
func RemoverUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Remover usuário"))
}
