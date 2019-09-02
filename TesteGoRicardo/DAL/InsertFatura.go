package DAL

import (
	"TesteGoRicardo/DataBase"
	"log"
	"strconv"
)

func InsertFatura(valor int, categoria string) bool {
	db := DataBase.DbConn()

	// Prepara a SQL e verifica errors
	insForm, err := db.Prepare("INSERT INTO fatura(categoria, valor) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do formulario com a SQL tratada e verifica errors
	_,err=insForm.Exec(categoria, valor)
	if err != nil {
		return false
	}

	// Exibe um log com os valores digitados no formulário
	log.Println("INSERT: Valor: " + strconv.Itoa(valor) + " | Categoria: " + categoria)

	// Encerra a conexão do dbConn()
	defer db.Close()
	return true

}