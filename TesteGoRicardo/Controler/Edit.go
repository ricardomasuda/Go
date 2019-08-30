package Controler

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
	"net/http"
)

// Função Edit, edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com banco de dados
	db := DataBase.DbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Model.Names{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		//Armazena os valores em variaveis
		var id int
		var name, email string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Name = name
		n.Email = email
	}

	// Mostra o template com formulário preenchido para edição
	tmpl.ExecuteTemplate(w, "Edit", n)

	// Fecha a conexão com o banco de dados
	defer db.Close()
}
