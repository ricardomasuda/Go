package main

import (
	"Fatura/Login"
	"Fatura/Model"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/fatura", Login.Fatura)

	http.HandleFunc("/funcao", Model.Funcao)
	http.HandleFunc("/login", Login.Login)

	http.Handle("/", http.FileServer(http.Dir("layout/dist/css/")))

	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}
