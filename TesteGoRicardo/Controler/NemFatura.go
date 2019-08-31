package Controler

import "net/http"

func NewFatura(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NewFatura", nil)
}
