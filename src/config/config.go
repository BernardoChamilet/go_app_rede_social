package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//pacote responsável por inicializar as variáveis de ambiente entre outra configurações

var (
	//ApiUrl representa url de onde api está rodando
	APIURL = ""
	//Porta representa a porta que o webapp está rodando
	Porta = 0
	//HashKey é utilizada para autenticar o cookie
	HashKey []byte
	//BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}
	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
