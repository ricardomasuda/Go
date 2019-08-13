package Model

import "html/template"

//Modelos armazena os modelos de pagina que serão executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("Pag/login.html"))


