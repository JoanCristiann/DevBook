package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"INSERT INTO usuarios(nome, username, email, senha) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(
		usuario.Nome,
		usuario.Username,
		usuario.Email,
		usuario.Senha,
	)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}
