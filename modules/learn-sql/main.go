package main

import (
	"database/sql"
	"learn-sql/userapi"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// connStr := "user=postgres password=gotraining dbname=postgres sslmode=disable"
	connStr := "postgres://postgres:gotraining@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(userapi.StartServer(db))
}
