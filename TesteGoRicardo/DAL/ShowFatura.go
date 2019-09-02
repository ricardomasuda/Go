package DAL

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
)

func ShowFatura(nId int) Model.Fatura {

	//abre conexão
	db := DataBase.DbConn()

	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT * FROM `fatura` WHERE `id_fatura`=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a strcut para ser utilizada no template
	n := Model.Fatura{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variaveis
		var id_fatura , valor int
		var categoria string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id_fatura, &valor, &categoria)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.IdFatura = id_fatura
		n.Categoria = categoria
		n.Valor = valor
	}
	// Fecha a conexão
	defer db.Close()

	return n
}
