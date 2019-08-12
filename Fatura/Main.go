package main

import (
	"Fatura/Login"
	"Fatura/Model"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})

	http.HandleFunc("/funcao", Model.Funcao)
	http.HandleFunc("/login", Login.Login)

	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}
