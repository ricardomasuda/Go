package Controler

import (
	"TesteGoRicardo/DataBase"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("tmpl/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := DataBase.DbConn()
	
	selDB, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "Index",nil )
}


