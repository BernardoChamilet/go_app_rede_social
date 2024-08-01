package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

//auxiliar do router para configurar as rotas (cada tipo de usuário terá um arquivo .go para definir as rotas (ex: usuarios da rede social tem usuarios.go))

// Rota representa todas as rotas da webaap
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca as rotas dentro do router, dependendo se estão autenticadas
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotaLogout)
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotasPublicacoes...)
	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	//configurando o app para encontrar a pasta assets e facilitando referência nos arquivos.html
	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
