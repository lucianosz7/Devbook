package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
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

func (repositorio usuarios) BuscarUsuarioID(usuarioID uint64) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", usuarioID)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Nick, &usuario.CriadoEm); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarAllUsuarios() ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query("select * from usuarios")
	if erro != nil {
		return []models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Nick, &usuario.Senha, &usuario.CriadoEm); erro != nil {
			return []models.Usuario{}, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) AtualizarUsuario(usuario models.Usuario) (models.Usuario, error) {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, email = ?, nick = ? where id = ?")
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, usuario.Nick, usuario.ID); erro != nil {
		return models.Usuario{}, erro
	}

	return usuario, nil
}

func (repositorio usuarios) DeletarUsuario(usuarioID uint64) (uint64, error) {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID); erro != nil {
		return 0, erro
	}

	return usuarioID, nil
}
