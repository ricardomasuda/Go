package main

import (
	"Fatura/DB"
	"Fatura/Login"
	"Fatura/Controladores"
	"fmt"
	"net/http"
)

func init()  {
	fmt.Println("vamos começar a subir o servidor")
	 _,err := DB.AbreConexao()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir banco de dados", err.Error())
		return
	}
}

func main() {


	http.HandleFunc("/fatura", Login.Fatura)

	http.HandleFunc("/funcao", Controladores.Funcao)
	http.HandleFunc("/login", Login.Login)
	http.HandleFunc("/cadastrar", Login.CadstrarFatura)


	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8282", nil)

}
