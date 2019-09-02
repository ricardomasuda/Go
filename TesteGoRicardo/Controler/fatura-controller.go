package Controler

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
	"log"
	"net/http"
	"strconv"
)

func EditFatura(w http.ResponseWriter, r *http.Request) {
	// Pega o ID do parametro da URL e  converte pra int
	nId,_ := strconv.Atoi(r.URL.Query().Get("id"))
	// Monta a struct para ser utilizada no template
	n := Model.Fatura{}
	//Busca a fatura no banco
	n = DAL.ShowFatura(nId)
	// Mostra o template com formulário preenchido para edição de fatura
	tmpl.ExecuteTemplate(w, "EditFatura", n)

}

func ShowFatura(w http.ResponseWriter, r *http.Request) {
	// Pega o ID do parametro da URL
	nId,_ :=strconv.Atoi(r.URL.Query().Get("id"))
	// Monta a strcut para ser utilizada no template
	n := Model.Fatura{}

	n=DAL.ShowFatura(nId)
	// Mostra o template
	tmpl.ExecuteTemplate(w, "ShowFatura", n)
}

// Função New apenas exibe o formulário para inserir novos dados
func NewFatura(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NewFatura", nil)
}

func InsertFatura(w http.ResponseWriter, r *http.Request) {
	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {
		// Pega os campos do formulário
		valor,_ := strconv.Atoi(r.FormValue("valor"))
		categoria := r.FormValue("categoria")
		status := r.FormValue("status")
		log.Print("---"+status+"---")

		DAL.InsertFatura(valor,categoria)
	}
	//Retorna a Lista de faturas
	http.Redirect(w, r, "listFatura", 301)
}

func ListFatura(w http.ResponseWriter, r *http.Request) {

	// Monta um array para guardar os valores da struct
	res := []Model.Fatura{}
	//Busca no banco todas as faturas
	res=DAL.ListFatura()

	tmpl.ExecuteTemplate(w, "ListFatura",res )
}

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
