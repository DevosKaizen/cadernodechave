package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Cria conexao com banco de dados e testa se tem erro
func ConectaCombancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=chaves_furg password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
