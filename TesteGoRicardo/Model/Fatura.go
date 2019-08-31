package Model

type Fatura struct {
	IdFatura  int    `db:"id_fatura" json:"id_fatura"`
	Valor     int    `db:"valor" json:"valor"`
	Categoria string `db:"categoria" json:"categoria"`
}
