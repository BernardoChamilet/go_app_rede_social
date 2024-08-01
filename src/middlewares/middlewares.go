package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

//fica no "meio" entre requisição e resposta. Funções que são aplicadas para todas rotas
//logo é útil para verificação do token
//usando essa abordagem a autenticação das rotas são verificas antes de colocalas no roteador(router). Eu poderia colocar todas e fazer controle de acesso só nos controllers
//esse proxima nessas funções significa que é pra passar para a próxima func aninhada. No final irá executar a função das rotas

// Logger escreve informações da requisição no terminal
func Logger(proxima http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		proxima(w, r)
	}
}

// Autenticar verifica a existência de cookies
func Autenticar(proxima http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//se n tiver cookies redireciona para /login
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		//se não segue para proxima func q devera ser a func da rota requisitada
		proxima(w, r)
	}
}
