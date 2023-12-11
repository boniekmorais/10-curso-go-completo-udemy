package controller

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Criar usuário.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	requestBody, error := io.ReadAll(r.Body)

	if error != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, error)
		return
	}

	var usuario model.Usuario

	if error = json.Unmarshal(requestBody, &usuario); error != nil {
		responses.Erro(w, http.StatusBadRequest, error)
		return
	}

	if erro := usuario.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, error := database.Conectar()

	if error != nil {
		responses.Erro(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	usuarioRepository := repository.NovoRepositorioUsuarios(db)
	usuario.ID, error = usuarioRepository.Criar(usuario)

	if error != nil {
		responses.Erro(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)

}

// Listar todos os usuários.
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {

	usuario := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := database.Conectar()

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	usuarioRepository := repository.NovoRepositorioUsuarios(db)

	usuarios, erro := usuarioRepository.Pesquisar(usuario)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)

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
