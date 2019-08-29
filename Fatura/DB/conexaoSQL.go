package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func AbreConexao()(db *sqlx.DB, err error){
	dbDriver := "mysql"
	dbUser := "extranet"
	dbPass := "Bm152612*"
	dbName := "goblog"
	db, err = sqlx.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro na conexao: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro no ping na conexao: ", err.Error())
		return
	}
	return
}

func GetDBConnection() (localdb *sqlx.DB, err error) {
	if db == nil {
		db, err = AbreConexao()
		if err != nil {
			log.Println("[GetDBConnection] Erro na conexao: ", err.Error())
			return
		}
	}
	err = db.Ping()
	if err != nil {
		log.Println("[GetDBConnection] Erro no ping na conexao: ", err.Error())
		return
	}
	localdb = db
	return
}