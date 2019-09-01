package Controler

import (
	"TesteGoRicardo/DAL"
	"log"
	"net/http"
	"strconv"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//pega o valor via GET
	nId,_ := strconv.Atoi(r.URL.Query().Get("id"))
	//retorno do DAL de delete User
	retorno := DAL.DeletarUser(nId)

	// Exibe um log com os valores digitados no form
	log.Println("DELETE")

	if retorno == true{
		// Retorna a HOME
		http.Redirect(w, r, "/", 301)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		name := r.FormValue("name")
		email := r.FormValue("email")
		id ,_ := strconv.Atoi(r.FormValue("uid"))

		//insert User
		DAL.UpdateUser(name, email , id);

		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: Name: " + name + " |E-mail: " + email)
	}
	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}
