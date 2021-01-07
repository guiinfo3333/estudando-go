package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	conexao := "user=postgres dbname=aluna_loja password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Nao deu erro")
	}
	return db
}
