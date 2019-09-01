package Controler

import (
	"TesteGoRicardo/DAL"
	"log"
	"net/http"
	"strconv"
)

func UpdateFatura(w http.ResponseWriter, r *http.Request) {

	var retorno bool
	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		valor , _:= strconv.Atoi(r.FormValue("valor"))
		categoria := r.FormValue("categoria")
		id,_ := strconv.Atoi(r.FormValue("uid"))

		//insere no banco as informação de fatura
		retorno = DAL.UpdateFatura(valor , categoria , id)

		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: categoria: " + categoria + " |faura: " + strconv.Itoa(valor))
	}
	// Retorna a HOME
	if retorno == true {
		http.Redirect(w, r, "/listFatura", 301)
	}
}

func DeleteFatura(w http.ResponseWriter, r *http.Request) {
	//receber valor do id
	nId,_ := strconv.Atoi(r.URL.Query().Get("id"))
	//retorno do DAL de deletar fatura
	retorno :=DAL.DeleteFatura(nId)
	// Exibe um log com os valores digitados no form
	log.Println("DELETE")
	if retorno == true{
		// Retorna a Home Fatura
		http.Redirect(w, r, "listFatura", 301)
	}

}
