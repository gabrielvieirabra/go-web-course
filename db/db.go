package db

import (
	"database/sql"
	_ "github.com/go-web-course/lib/pq"
	"log"
)

func ConectaBD() *sql.DB {
	connStr := "user=postgres dbname=alura_loja password=2eY7ED41 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
