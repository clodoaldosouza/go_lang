package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

/*
	https://pkg.go.dev/
	need to run once:  go get github.com/lib/pq
	the underscore before the "github.com/lib/pq" mens a Runtime import
*/

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
