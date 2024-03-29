package DAL

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
)

func ListFatura()  []Model.Fatura {

	db := DataBase.DbConn()

	selDB, err := db.Query("SELECT * FROM `fatura` WHERE 1")
	if err != nil {
		panic("[ListFatura]Erro ao buscar informações no banco"+err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Model.Fatura{}

	// Monta um array para guardar os valores da struct
	res := []Model.Fatura{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var idFatura,valor int
		var categoria,status string


		// Faz o Scan do SELECT
		err = selDB.Scan(&idFatura, &categoria, &valor,&status)
		if err != nil {
			panic("[ListFatura]Erro ao fazer o Scan do SELECT"+ err.Error())
		}

		// Envia os resultados para a struct
		n.IdFatura = idFatura
		n.Categoria = categoria
		n.Valor = valor
		n.Status = status

		// Junta a Struct com Array de Struct Fatura
		res = append(res, n)
	}
	return res
}
