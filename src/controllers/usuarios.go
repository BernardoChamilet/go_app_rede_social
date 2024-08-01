package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

//pacote responsável por receber requisições e fazer devidas ações antes de interagir com a api

// CriarUsuario chama a api para cadastrar um usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	//pegando dados do arquivo js
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	//perceba que se der um erro ao cadastrar, a api retornara uma resposta igual
	//portanto erro continuará sendo nil e será necessária ver se o statuscode da resposta equivale a um erro
	//isso é necessário para saber porque deu erro. Para avisar o usuário o motivo
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	//mandando resposta da requisição que vai poder ser lida no js
	//AQui em baixo tem um código para fazer um json vazio ao invés de mandar nil
	//isso foi feito para corrigir um erro que tava dando no ajax
	jsonVazio := map[string]string{"status": "success"}
	jsonVazioR, err := json.Marshal(jsonVazio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respostas.JSON(w, response.StatusCode, jsonVazioR)
}

// PararDeSeguir para de seguir um usuário com auxilio da api
func PararDeSeguir(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// Seguir segue um usuário com auxilio da api
func Seguir(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// EditarUsuario edita os dados de um usuário com auxilio da api
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	//pegando dados do arquivo js
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//pegando id do cookie
	cookie, _ := cookies.Ler(r)
	logadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, logadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)

}

// EditarSenha edita a senha de um usuário com auxilio da api
func EditarSenha(w http.ResponseWriter, r *http.Request) {
	//pegando dados do arquivo js
	r.ParseForm()
	senhas, erro := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova":  r.FormValue("nova"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//pegando id do cookie
	cookie, _ := cookies.Ler(r)
	logadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.APIURL, logadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

// DeletarUsuario deleta um usuário com auxilio da api
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	//pegando id do cookie
	cookie, _ := cookies.Ler(r)
	logadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, logadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}
