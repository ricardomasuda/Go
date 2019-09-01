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
	http.HandleFunc("/showFatura", Controler.ShowFatura)
	http.HandleFunc("/new", Controler.New)
	http.HandleFunc("/newFatura", Controler.NewFatura)
	http.HandleFunc("/edit", Controler.Edit)
	http.HandleFunc("/editFatura", Controler.EditFatura)
	http.HandleFunc("/listFatura",Controler.ListFatura)

	// Ações
	http.HandleFunc("/insert", Controler.Insert)
	http.HandleFunc("/insertFatura", Controler.InsertFatura)
	http.HandleFunc("/update", Controler.UpdateUser)
	http.HandleFunc("/updateFatura", Controler.UpdateFatura)
	http.HandleFunc("/delete", Controler.DeleteUser)
	http.HandleFunc("/deleteFatura", Controler.DeleteFatura)

	// Inicia o servidor na porta 9000/ link localhost:9000
	http.ListenAndServe(":9001", nil)
}
