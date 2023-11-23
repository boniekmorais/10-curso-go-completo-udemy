package controller

import "net/http"

// Criar usuário.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar usuário"))
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
