package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarAllUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuarioID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controller.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}/deixar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controller.UnfollowUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controller.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
