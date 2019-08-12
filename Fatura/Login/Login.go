package Login

import (
	"Fatura/Model"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	r.ParseForm()
	// logic part of log in
	fmt.Println("username:", r.Form["Login"])
	fmt.Println("password:", r.Form["Senha"])
	if err := Model.Login.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Login] Erro na execucao do modelo: ", err.Error())
	}
}
