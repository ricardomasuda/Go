package Model

import (
	"fmt"
	"net/http"
)

func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando função num pacote")
}
