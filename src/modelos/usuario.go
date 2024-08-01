package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario representa usuario da rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoem"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto busca todas as informações de um usuário usando a api com concorrência
func BuscarUsuarioCompleto(usuarioId uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosUsuario(canalUsuario, usuarioId, r)
	go BuscarSeguidores(canalSeguidores, usuarioId, r)
	go BuscarSeguindo(canalSeguindo, usuarioId, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioId, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	//ATENÇÃO: aqui o go considera um slice vazio com nil
	//iterando 4 vezes para ter certeza que todos os canais receberam algo
	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			//se deu erro
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar dados do usuário")
			}

			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			//se deu erro
			if seguidoresCarregados == nil && reflect.TypeOf(seguidoresCarregados).Kind() != reflect.Slice {
				return Usuario{}, errors.New("erro ao buscar dados de seguidores")
			}

			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			//se deu erroS
			if seguindoCarregados == nil && reflect.TypeOf(seguindoCarregados).Kind() != reflect.Slice {
				return Usuario{}, errors.New("erro ao buscar dados de seguindo")
			}

			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-canalPublicacoes:
			//se deu erro
			if publicacoesCarregadas == nil && reflect.TypeOf(publicacoesCarregadas).Kind() != reflect.Slice {
				return Usuario{}, errors.New("erro ao buscar dados de publicações")
			}

			publicacoes = publicacoesCarregadas
		}
	}
	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

// BuscarDadosUsuario chama a api para buscar dados base do usuário
func BuscarDadosUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()
	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}
	canal <- usuario

}

// BuscarSeguidores chama a api para buscar dados dos seguidores de um usuário
func BuscarSeguidores(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()
	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}
	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguidores
}

// BuscarSeguindo chama a api para buscar dados de quem um usuário segue
func BuscarSeguindo(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()
	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}
	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguindo
}

// BuscarPublicacoes chama a api para buscar todas as publicações de um usuário
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	//chamando api
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()
	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}
	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}
	canal <- publicacoes
}
