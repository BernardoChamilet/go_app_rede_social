package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// configurar utiliza as variáveis hashkey e blockkey para criar um securecookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar registra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {
	//pega dados do usuario logado
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}
	//codifica eles com o cookie criado
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}
	//salva de fato
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

// Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	//pegando cookie codificado
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}
	//descodificando e passando para um map
	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}
	return valores, nil

}

// Deletar deleta o valor de dentro do cookie do navegador
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})

}
