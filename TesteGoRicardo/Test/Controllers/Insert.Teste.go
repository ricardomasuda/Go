package Controllers

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
)

func testeInsert() bool{

	var check bool
	//1 setup exemplo : Criar e salvar  usuario fake no banco
	//para o insert não é preciso

	//2 exercise exemplo : Chamar o controller
	name := "RICARDO"
	email:= "mricardo1611@gmail.com"
	id:=DAL.InsertUser(name , email)

	//3 verify : checa se o retorno é o esperado
	n := Model.Names{}
	n = DAL.ShowUser(id)
	if n.Name == name && n.Email==email {
		check = true
	}  else{
		check = false
	}

	//4 teardown : limpa os dados fakes
	DAL.DeletarUser(id)

	//retorna o dado do teste
	return  check
}
