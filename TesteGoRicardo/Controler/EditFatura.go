package Controler

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
	"net/http"
)

func EditFatura(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com banco de dados
	db := DataBase.DbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM `fatura` WHERE `id_fatura` =?", nId)
		if err != nil {
		panic("[EditarFatura]Erro ao executar o SQL "+err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Model.Fatura{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		//Armazena os valores em variaveis
		var id_fatura , valor int
		var categoria string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id_fatura, &categoria, &valor)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.IdFatura = id_fatura
		n.Categoria = categoria
		n.Valor = valor
	}

	// Mostra o template com formulário preenchido para edição de fatura
	tmpl.ExecuteTemplate(w, "EditFatura", n)

	// Fecha a conexão com o banco de dados
	defer db.Close()
}