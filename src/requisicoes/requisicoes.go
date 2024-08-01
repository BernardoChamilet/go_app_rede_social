package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

//Pacote responsável por ajudar a fazer as requisições que requerem autenticação pra api

// FazerRequisicaoComAutenticacao é utilizada para colocar o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	//criando requisição que será feita a api
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}
	//pegando token do cookie e colocando no header da requisição
	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	//craindo um cliente para fazer a requisição e obter a resposta da api
	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
