package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID            uint64    `json:"id,omitempty"`
	Titulo        string    `json:"titulo,omitempty"`
	Conteudo      string    `json:"conteudo,omitempty"`
	AutorID       uint64    `json:"autorId,omitempty"`
	AutorUsername string    `json:"autorUsername,omitempty"`
	Likes         uint64    `json:"likes"`
	CriadaEm      time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o titulo é obrigatório e não pode estar em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("o conteudo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
