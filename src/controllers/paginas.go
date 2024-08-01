package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaDeLogin renderiza página de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario renderiza página de cadastro de usuários
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal carrega página home com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	//chamando função de BuscarPublicações da api
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	//guardando publis numa variavel
	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//pegando cookie para saber quem está logado
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//mandando dados para home.html
	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioId,
	})
}

// CarregarPaginaDeEditarPublicacao renderiza página de editar ppublicações
func CarregarPaginaDeEditarPublicacao(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	//mandando dados para editar-publicacao.html
	var publicacao modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "editar-publicacao.html", publicacao)
}

// CarregarPaginaDeUsuarios renderiza página de usuarios que atendem filtro passado
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	//pegando variaveis dos parametros da url
	nomeOUnick := strings.ToLower(r.URL.Query().Get("usuario"))
	//chamando api
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOUnick)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	//Criando slice de usuários para guardar resposta da api e enviar pro frontend
	var usuarios []modelos.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario renderiza página de perfil de um usuário
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//vendo se o usuário logado é o usuário que a página está sendo carregada
	//nesse caso tem que carregar a página /perfil e não /usuarios/id
	//E caso não, vendo se o usuário logado segue o usuário que a página está sendo carregada
	cookie, _ := cookies.Ler(r)
	logadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	if usuarioId == logadoID {
		http.Redirect(w, r, "/perfil", http.StatusFound)
		return
	}
	//chamando função que acessa várias rotas da api com concorrência para obter todas informações de um usuário
	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioId, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario  modelos.Usuario
		LogadoID uint64
	}{
		Usuario:  usuario,
		LogadoID: logadoID,
	})
}

// CarregarPerfilDoLogado renderiza página de perfil do usuário logado
func CarregarPerfilDoLogado(w http.ResponseWriter, r *http.Request) {
	//pegando id do usuario logado
	cookie, _ := cookies.Ler(r)
	logadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//chamando função que acessa várias rotas da api com concorrência para obter todas informações de um usuário
	usuario, erro := modelos.BuscarUsuarioCompleto(logadoID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	utils.ExecutarTemplate(w, "perfil.html", usuario)
}

// CarregarEditarUsuario renderiza página de edição de dados do usuário logado
func CarregarEditarUsuario(w http.ResponseWriter, r *http.Request) {
	//pegando id do usuario logado
	cookie, _ := cookies.Ler(r)
	logadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//usando api para pegar dados
	canal := make(chan modelos.Usuario)
	go modelos.BuscarDadosUsuario(canal, logadoID, r)
	usuario := <-canal
	if usuario.ID == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "erro ao buscar dados do usuário"})
		return
	}
	utils.ExecutarTemplate(w, "editar-usuario.html", usuario)
}

// CarregarEditarSenha renderiza página de edição de senha do usuário logado
func CarregarEditarSenha(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "editar-senha.html", nil)
}
