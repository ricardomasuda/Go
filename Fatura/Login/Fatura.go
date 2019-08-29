package Login

import (
	"Fatura/Controladores"
	"Fatura/DB"
	"Fatura/Model"
	"fmt"
	"log"
	"net/http"
)


func Fatura(w http.ResponseWriter, r *http.Request) {

	fatura := Model.Fatura{}

	fmt.Println("method:", r.Method)
	r.ParseForm()
	// logic part of log in
	db, err := DB.GetDBConnection()
	if err != nil {
		log.Println("[Local] Erro na conexao: ", err.Error())
		return
	}
	sql :=`INSERT INTO  user ( nome ,  senha ) VALUES ('Ricardo','Senha')`
	resultado, err := db.Exec(sql)
	if err != nil {
		fmt.Println("[local] Erro na inclusao do log: ", sql, " - ", err.Error())
	}
	println(resultado)

	if err := Controladores.Fatura.ExecuteTemplate(w, "fatura.html", fatura); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[login] Erro na execucao do modelo: ", err.Error())
	}
}