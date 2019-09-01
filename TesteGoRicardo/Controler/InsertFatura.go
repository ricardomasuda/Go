package Controler

import (
	"TesteGoRicardo/DataBase"
	"log"
	"net/http"
)

func InsertFatura(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		valor := r.FormValue("valor")
		categoria := r.FormValue("categoria")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO fatura(categoria, valor) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulario com a SQL tratada e verifica errors
		result , err := insForm.Exec(categoria, valor)


		log.Println(result)
		// Exibe um log com os valores digitados no formulário
		log.Println("INSERT: Valor: " + valor + " | Categoria: " + categoria)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	//Retorna a Lista de faturas
	http.Redirect(w, r, "listFatura", 301)
}
