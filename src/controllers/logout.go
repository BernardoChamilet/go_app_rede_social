package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout tira cookie de autenticação do navegador
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
