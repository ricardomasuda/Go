package Controler

import (
	"TesteGoRicardo/DataBase"
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {

	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("uid")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("UPDATE names SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulário com a SQL tratada e verifica erros
		insForm.Exec(name, email, id)

		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: Name: " + name + " |E-mail: " + email)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}