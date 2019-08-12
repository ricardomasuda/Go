package Model

import "html/template"

//Modelos armazena os modelos de pagina que serão executados pelos manipuladores
var Login = template.Must(template.ParseFiles("Pag/login.html"))

//Modelos armazena os modelos de pagina que serão executados pelos manipuladores
var Fatura = template.Must(template.ParseFiles("Pag/fatura.html"))