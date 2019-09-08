package Controllers

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
)

/*
4 fases de teste*/

func testeShow() bool{

	var check bool
	name := "RICARDO"
	email:= "mricardo1611@gmail.com"
	//1 setup exemplo : Criar e salvar  usuario fake no banco
	id:=DAL.InsertUser(name , email)
	user:=Model.Names{}

	//2 exercise exemplo : Chamar o controller
	user=DAL.ShowUser(id)

	//3 verify : checa se o retorno é o esperado
	if user.Id == id {
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