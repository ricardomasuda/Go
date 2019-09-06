package main

import (
	"TesteGoRicardo/Controler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:9000")
	rtr := mux.NewRouter()
	// Gerencia as URLs
	//http.HandleFunc("/", Controler.Index)
	rtr.HandleFunc("/", Controler.Index)
	rtr.HandleFunc("/show", Controler.Show)
	rtr.HandleFunc("/showFatura", Controler.ShowFatura)
	rtr.HandleFunc("/newUser", Controler.New)
	rtr.HandleFunc("/newFatura", Controler.NewFatura)
	rtr.HandleFunc("/edit", Controler.Edit)
	rtr.HandleFunc("/editFatura", Controler.EditFatura)
	rtr.HandleFunc("/listaFatura", Controler.ListFatura)

	//// Ações
	rtr.HandleFunc("/insert", Controler.Insert)
	rtr.HandleFunc("/insertFatura", Controler.InsertFatura)
	rtr.HandleFunc("/update", Controler.UpdateUser)
	rtr.HandleFunc("/updateFatura", Controler.UpdateFatura)
	rtr.HandleFunc("/deleteUser", Controler.DeleteUser)
	rtr.HandleFunc("/deleteFatura", Controler.DeleteFatura)

	//bootstrap
	rtr.PathPrefix("/layout/").Handler(http.StripPrefix("/layout/", http.FileServer(http.Dir("./layout"))))

	http.Handle("/",rtr)
	// Inicia o servidor na porta 9000/ link localhost:9000
	http.ListenAndServe(":9001", nil)
}
