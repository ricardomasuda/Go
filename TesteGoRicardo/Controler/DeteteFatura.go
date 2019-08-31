package Controler

import (
	"TesteGoRicardo/DataBase"
	"log"
	"net/http"
)

func DeleteFatura(w http.ResponseWriter, r *http.Request) {

	println("entrou em delete")
	// Abre conexão com banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	nId := r.URL.Query().Get("id")

	// Prepara a SQL e verifica errors
	delForm, err := db.Prepare("DELETE FROM fatura WHERE id_fatura=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do form com a SQL tratada e verifica errors
	delForm.Exec(nId)

	// Exibe um log com os valores digitados no form
	log.Println("DELETE :"+nId)

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "listFatura", 301)
}