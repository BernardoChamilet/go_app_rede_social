package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarPublicacao cria uma publicação com ajuda da api
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	//pegando dados do arquivo js
	r.ParseForm()
	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
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

// CurtirPublicacao curti uma publicação com ajuda da api
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.APIURL, publicacaoId)
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
	//mandando resposta da requisição que vai poder ser lida no js
	respostas.JSON(w, response.StatusCode, nil)
}

// DescurtirPublicacao descurti uma publicação com ajuda da api
func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/publicacoes/%d/descurtir", config.APIURL, publicacaoId)
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
	//mandando resposta da requisição que vai poder ser lida no js
	respostas.JSON(w, response.StatusCode, nil)
}

// EditarPublicacao edita uma publicação com ajuda da api
func EditarPublicacao(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//pegando dados do arquivo js
	r.ParseForm()
	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//chamando api
	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	//mandando resposta da requisição que vai poder ser lida no js
	respostas.JSON(w, response.StatusCode, nil)
}

// DeletarPublicacao deleta uma publicação com ajuda da api
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	//pegando parametros passados pelo js
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoId)
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
	//mandando resposta da requisição que vai poder ser lida no js
	respostas.JSON(w, response.StatusCode, nil)
}
