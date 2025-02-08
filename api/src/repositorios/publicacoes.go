package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(
		publicacao.Titulo,
		publicacao.Conteudo,
		publicacao.AutorID,
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

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, err := repositorio.db.Query(`
		SELECT P.*, U.username FROM publicacoes P
		INNER JOIN usuarios U ON U.id = P.autor_id
		WHERE P.id = ?`,
		publicacaoID,
	)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao models.Publicacao
	if linha.Next() {
		if err = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Likes,
			&publicacao.CriadaEm,
			&publicacao.AutorUsername,
		); err != nil {
			return models.Publicacao{}, err
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
	SELECT DISTINCT P.*, U.username FROM publicacoes P
	INNER JOIN usuarios U ON U.id = P.autor_id
	INNER JOIN seguidores S ON P.autor_id = S.usuario_id
	WHERE U.id = ? OR S.seguidor_id = ?
	ORDER BY P.criadaEm DESC`,
		usuarioID,
		usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Likes,
			&publicacao.CriadaEm,
			&publicacao.AutorUsername,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, err := repositorio.db.Prepare(
		"UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		publicacao.Titulo,
		publicacao.Conteudo,
		publicacaoID,
	); err != nil {
		return err
	}

	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM publicacoes WHERE ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
		SELECT P.*, U.username FROM publicacoes P
		INNER JOIN usuarios U ON U.id = P.autor_id
		WHERE P.autor_id = ?
		ORDER BY P.criadaEm DESC`,
		usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Likes,
			&publicacao.CriadaEm,
			&publicacao.AutorUsername,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}
