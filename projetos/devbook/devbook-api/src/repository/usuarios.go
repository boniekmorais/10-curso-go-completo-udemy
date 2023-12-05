package repository

import (
	"api/src/model"
	"database/sql"
)

// Representa um repositório de usuários.
type Usuarios struct {
	db *sql.DB
}

// Cria um repositório de usuários.
func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Insere usuário no banco de dados.
func (u Usuarios) Criar(usuario model.Usuario) (uint64, error) {
	return 0, nil
}
