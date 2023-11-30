package model

import "time"

// Representa um usuário do sistema.
type Usuario struct {
	ID     uint64    `json:"id,omitempty"`
	Nome   string    `json:"nome,omitempty"`
	Nick   string    `json:"nick,omitempty"`
	Email  string    `json:"email,omitempty"`
	Senha  string    `json:"senha,omitempty"`
	Criado time.Time `json:"criado,omitempty"`
}