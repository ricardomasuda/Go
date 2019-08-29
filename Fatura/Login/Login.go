package Login

import (
	"Fatura/DB"
	"Fatura/Controladores"
	"fmt"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	r.ParseForm()
	// logic part of log in
	fmt.Println("username:", r.Form["Login"])
	fmt.Println("password:", r.Form["Senha"])
	Login := r.Form["Login"]
	Senha := r.Form["Senha"]
	fmt.Println(Login)
	fmt.Println(Senha)
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
	fmt.Println(resultado);
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[local] Erro ao pegar o numero de linhas afetadas pelo comando: ", sql, " - ", err.Error())
	}
	fmt.Println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")
	if err := Controladores.Modelos.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
	}
}
