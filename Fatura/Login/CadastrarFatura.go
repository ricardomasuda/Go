package Login

import (
	"Fatura/Controladores"
	"Fatura/DB"
	"fmt"
	"log"
	"net/http"
)

func CadstrarFatura(w http.ResponseWriter, r *http.Request) {

	if err := Controladores.CadastrarFatura.ExecuteTemplate(w, "cadastrarFatura.html", nil); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
	}

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

	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[local] Erro ao pegar o numero de linhas afetadas pelo comando: ", sql, " - ", err.Error())
	}

	fmt.Println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")

}
