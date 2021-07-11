package db

import (
	"database/sql"

	_ "github.com/lib/pq" //Nao iremos usar a lib o tempo interito, então o _ serve para isso
)

func ConectaComBD() *sql.DB {
	conexao := "user=postgres dbname=camila_loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
