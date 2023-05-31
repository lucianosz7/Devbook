package controller

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(body, &publicacao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.CriarPub(publicacao)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPublicacao(publicacaoId)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusOK, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	var publicacoes []models.Publicacao
	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacoes, erro = repositorio.BuscarPublicacoes(usuarioID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, publicacoes)
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	pubSalva, erro := repositorio.BuscarPublicacao(publicacaoId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if pubSalva.AutorID != usuarioID {
		response.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua"))
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(body, &publicacao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Preparar(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarPub(publicacaoId, publicacao); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := auth.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	pubSalva, erro := repositorio.BuscarPublicacao(publicacaoId)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if pubSalva.AutorID != usuarioID {
		response.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar uma publicação que não seja sua"))
		return
	}

	if erro = repositorio.Deletar(publicacaoId); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	if erro = repositorio.Curtir(publicacaoID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	if erro = repositorio.Descurtir(publicacaoID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
