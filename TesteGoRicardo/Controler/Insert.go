package Controler

import (
	"TesteGoRicardo/DataBase"
	"log"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		name := r.FormValue("name")
		email := r.FormValue("email")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO names(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulario com a SQL tratada e verifica errors
		insForm.Exec(name, email)

		// Exibe um log com os valores digitados no formulário
		log.Println("INSERT: Name: " + name + " | E-mail: " + email)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	//Retorna a HOME
	http.Redirect(w, r, "/", 301)
}