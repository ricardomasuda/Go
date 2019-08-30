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
	http.HandleFunc("/show", Controler.Show)
	http.HandleFunc("/new", Controler.New)
	http.HandleFunc("/edit", Controler.Edit)

	// Ações
	http.HandleFunc("/insert", Controler.Insert)
	http.HandleFunc("/update", Controler.Update)
	http.HandleFunc("/delete", Controler.Delete)

	// Inicia o servidor na porta 9000/ link localhost:9000
	http.ListenAndServe(":9000", nil)
}
