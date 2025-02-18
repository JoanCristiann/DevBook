package models

import "time"

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID            uint64    `json:"id,omitempty"`
	Titulo        string    `json:"titulo,omitempty"`
	Conteudo      string    `json:"conteudo,omitempty"`
	AutorID       uint64    `json:"autorId,omitempty"`
	AutorUsername string    `json:"autorUsername,omitempty"`
	Likes         uint64    `json:"likes"`
	CriadaEm      time.Time `json:"criadaEm,omitempty"`
}
