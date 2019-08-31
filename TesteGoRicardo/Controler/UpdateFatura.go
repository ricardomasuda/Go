package Controler

import (
	"TesteGoRicardo/DataBase"
	"log"
	"net/http"
)

func UpdateFatura(w http.ResponseWriter, r *http.Request) {

	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		valor := r.FormValue("valor")
		categoria := r.FormValue("categoria")
		id := r.FormValue("uid")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("UPDATE fatura SET valor=?, categoria=? WHERE id_fatura=?")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulário com a SQL tratada e verifica erros
		insForm.Exec(valor, categoria, id)

		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: categoria: " + categoria + " |faura: " + valor)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/listFatura", 301)
}
