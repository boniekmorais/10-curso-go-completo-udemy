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
func (usuarioRepository Usuarios) Criar(usuario model.Usuario) (uint64, error) {

	statement, erro := usuarioRepository.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
