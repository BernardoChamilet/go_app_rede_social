package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	//carregando variáveis de ambiente
	config.Carregar()
	//configurando cookies
	cookies.Configurar()
	//carregando páginas html
	utils.CarregarTemplates()
	//gerando rotas
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d \n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
