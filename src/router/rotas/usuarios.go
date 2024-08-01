package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/buscar-usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguir,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Seguir,
		RequerAutenticacao: true,
	},
	{
		URI:                "/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoLogado,
		RequerAutenticacao: true,
	},
	{
		URI:                "/editar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarEditarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/editar-usuario",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/editar-senha",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarEditarSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/editar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.EditarSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/deletar-usuario",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
}
