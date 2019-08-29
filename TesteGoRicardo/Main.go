package main

import (
	"TesteGoRicardo/Controler"
	"log"
	"net/http"
)



func main() {
	log.Println("Server started on: http://localhost:9000")

	// Gerencia as URLs
	http.HandleFunc("/", Controler.Index)
	//http.HandleFunc("/show", Show)
	//http.HandleFunc("/new", New)
	//http.HandleFunc("/edit", Edit)

	// Ações
	http.HandleFunc("/insert", Controler.Insert)
	//http.HandleFunc("/update", Update)
	//http.HandleFunc("/delete", Delete)

	// Inicia o servidor na porta 9000/ link localhost:9000
	http.ListenAndServe(":9000", nil)
}
