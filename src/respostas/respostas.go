package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//pacote responsavel por padronizar as resposta que o app dará ao frontend

// Erro representa a resposta de statuscode de erro vinda com sucesso da api
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em json a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	//transformando em json
	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// TratarStatusCodeDeErro trata o statuscode de erro (400 >=)
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	//aqui eu obtenho o conteudo do erro, para mandar pro js pra ele poder avisar o usuário o motivo do erro
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
