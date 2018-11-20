package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func printAll(db *sql.DB) {
	//Query with no condition
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var (
			id                         int
			firstName, lastName, email string
		)
		rows.Scan(&id, &firstName, &lastName, &email)
		fmt.Println(id, firstName, lastName, email)
	}
}

func printByID(db *sql.DB, id int) {
	//Query with condition
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	var firstName, lastName, email string
	row.Scan(&id, &firstName, &lastName, &email)
	fmt.Println(id, firstName, lastName, email)
}

func main() {
	// connStr := "user=postgres password=gotraining dbname=postgres sslmode=disable"
	connStr := "postgres://postgres:gotraining@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// CREATE TABLE users(id SERIAL PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT);

	// Insert
	_, err = db.Exec("INSERT INTO users(first_name, last_name, email) VALUES ($1,$2,$3)", "Lord", "Gift", "test@test")
	if err != nil {
		log.Fatal(err)
	}

	printAll(db)

	// Update
	_, err = db.Exec("UPDATE users SET email = $1 WHERE id = $2", "updated@test", 1)
	if err != nil {
		log.Fatal(err)
	}
	printByID(db, 1)

	// DELETE
	_, err = db.Exec("DELETE FROM users WHERE id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}
	printByID(db, 1)

}
