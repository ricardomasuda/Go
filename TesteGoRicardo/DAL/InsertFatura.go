package DAL

import (
	"TesteGoRicardo/DataBase"
	"log"
	"strconv"
)

func InsertFatura(valor int, categoria string , status string) int {
	db := DataBase.DbConn()

	// Prepara a SQL e verifica errors
	insForm, err := db.Prepare("INSERT INTO fatura(categoria, valor,status) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do formulario com a SQL tratada e verifica errors
	id,err :=insForm.Exec(categoria, valor,status)
	if err != nil {
		return 0
	}

	LastInsertId, err := id.LastInsertId()
	// Exibe um log com os valores digitados no formulário
	log.Println("INSERT: ID " + strconv.FormatInt(LastInsertId, 10) + " Valor: " + strconv.Itoa(valor) + " | Categoria: " + categoria)
	// Encerra a conexão do dbConn()
	defer db.Close()
	return int(LastInsertId)

}