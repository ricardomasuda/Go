package Model

type Fatura struct {
	IdFatura int `json:"userID" db:"id_fatura"`
	Valor string `json:"Nome" db:"valor"`
	Categoria string `json:"Senha" db:"categoria"`
}