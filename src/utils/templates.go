package utils

import (
	"net/http"
	"text/template"
)

//Pacote que contém utilidades

var templates *template.Template

// CarregarTemplates insere os templates html na variável templates
func CarregarTemplates() {
	//aqui estou falando que os templates vão estar na pasta views e que são arquivos .html
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTemplate renderiza uma página html
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
