package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacoes,
		RequerAutenticacao: true,
	},

	{
		URI:                "/publicacoes/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacao,
		RequerAutenticacao: true,
	},

	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarPublicacao,
		RequerAutenticacao: true,
	},

	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controller.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controller.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
