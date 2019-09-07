package Controllers

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
)

/*
4 fases de teste*/

func testeShow() bool{

	var check bool
	//1 setup exemplo : Criar e salvar  usuario fake no banco
	name := "RICARDO"
	email:= "mricardo1611@gmail.com"
	id:=DAL.InsertUser(name , email)
	fatura:=Model.Fatura{}

	//2 exercise exemplo : Chamar o controller
	fatura=DAL.ShowFatura(id);

	//3 verify : checa se o retorno é o esperado
	if fatura.IdFatura == id {
		check = true
	} else {
		check = false
	}
	//4 teardown : limpa os dados fakes
	DAL.DeletarUser(id)

	//retorna o dado do teste
	return  check
}
/*


4 teardown : limpa os dados fakes
*/