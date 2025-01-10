package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório")
	}

	if usuario.Username == "" {
		return errors.New("o username é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("a senha é obrigatória")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Username = strings.TrimSpace(usuario.Username)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
