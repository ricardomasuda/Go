package Controladores

import "html/template"

//Modelos armazena os modelos de pagina que serão executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("Pag/login.html"))

var Fatura = template.Must(template.ParseFiles("Pag/fatura.html"))

var CadastrarFatura = template.Must(template.ParseFiles("Pag/cadastrarFatura.html"))

