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

func (repositorio Usuarios) Buscar(nomeOuUsername string) ([]models.Usuario, error) {
	nomeOuUsername = "%" + nomeOuUsername + "%"

	linhas, err := repositorio.db.Query(
		"SELECT id, nome, username, email, criadoEm FROM usuarios WHERE nome LIKE ? OR username LIKE ?",
		nomeOuUsername, nomeOuUsername)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Username,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"SELECT id, nome, username, email, criadoEm FROM usuarios WHERE id = ?", ID)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Username,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, username = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		usuario.Nome,
		usuario.Username,
		usuario.Email,
		ID,
	); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) Excluir(ID uint64) error {
	statement, err := repositorio.db.Prepare(
		"DELETE FROM usuarios WHERE id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, err := repositorio.db.Query(
		"SELECT id, senha FROM usuarios WHERE email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare(
		"INSERT IGNORE INTO seguidores(usuario_id, seguidor_id) VALUES(?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare(
		"DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		SELECT u.id, u.nome, u.username, u.email, u.criadoEm
		FROM usuarios u
		INNER JOIN seguidores s ON s.seguidor_id = u.id
		WHERE s.usuario_id = ?
	`, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Username,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		SELECT u.id, u.nome, u.username, u.email, u.criadoEm
		FROM usuarios u
		INNER JOIN seguidores s ON s.usuario_id = u.id
		WHERE s.seguidor_id = ?
	`, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Username,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, err := repositorio.db.Query(
		"SELECT senha FROM usuarios WHERE id = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer linha.Close()

	var senha string

	if linha.Next() {
		if err = linha.Scan(&senha); err != nil {
			return "", err
		}
	}

	return senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, err := repositorio.db.Prepare(
		"UPDATE usuarios SET senha = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(senha, usuarioID); err != nil {
		return err
	}

	return nil
}
