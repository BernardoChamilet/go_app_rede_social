package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

//pacote responsável por criar as rotas do app, utilizando da ajuda do pacote rotas (não sei porque não fazer tudo aqui)

// Gerar gera um roteador com todas rotas configuradas
func Gerar() *mux.Router {
	return rotas.Configurar(mux.NewRouter())
}
