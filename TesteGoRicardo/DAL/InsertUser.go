package DAL

import (
	"TesteGoRicardo/DataBase"
	"log"
)

func InsertUser(name string , email string )(bool) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Prepara a SQL e verifica errors
	insForm, err := db.Prepare("INSERT INTO names(name, email) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do formulario com a SQL tratada e verifica errors
	_,err=insForm.Exec(name, email)
	if err != nil {
		return false
	}

	// Exibe um log com os valores digitados no formulário
	log.Println("INSERT: Name: " + name + " | E-mail: " + email)

	// Encerra a conexão do dbConn()
	defer db.Close()
	return true
}