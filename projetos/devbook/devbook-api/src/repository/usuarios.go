package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

// Pesquisa usuários com base no filtro do nome ou nick.
func (usuarioRepository Usuarios) Pesquisar(usuario string) ([]model.Usuario, error) {

	usuario = fmt.Sprintf("%%%s%%", usuario) // Ex: %<nome>%

	registros, erro := usuarioRepository.db.Query(
		"SELECT id, nome, nick, email, criado FROM usuarios u WHERE u.nome LIKE ? OR u.nick LIKE ?",
		usuario, usuario,
	)

	if erro != nil {
		return nil, erro
	}

	defer registros.Close()

	var usuarios []model.Usuario

	for registros.Next() {

		var user model.Usuario

		if erro = registros.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.Criado,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, user)
	}

	return usuarios, nil
}
