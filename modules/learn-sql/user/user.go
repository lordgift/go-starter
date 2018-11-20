package user

import (
	"database/sql"
	"log"
)

type Service struct {
	DB *sql.DB
}

func (s *Service) FindByID(id int) (*User, error) {
	stmt := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"
	row := s.DB.QueryRow(stmt, id)
	var u User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

var db *sql.DB

func ConnectDB() {
	var err error
	connStr := "postgres://tpxbvdny:3xQCeyAUtP6Hq5uqgPCTrI54cfRUqzTe@elmer.db.elephantsql.com:5432/tpxbvdny"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func Insert(u *User) error {
	stmt := `INSERT INTO users(first_name, last_name, email)
		 values ($1, $2, $3) RETURNING id`
	row := db.QueryRow(stmt, u.FirstName, u.LastName, u.Email)
	err := row.Scan(&u.ID)

	return err
}

func FindByID(id int) (*User, error) {
	stmt := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"
	row := db.QueryRow(stmt, id)
	var u User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func All() ([]User, error) {
	stmt := "SELECT id, first_name, last_name, email FROM users ORDER BY id DESC"
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	var us []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

func Update(u *User) error {
	stmt := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4"
	_, err := db.Exec(stmt, u.FirstName, u.LastName, u.Email, u.ID)
	return err
}

func Delete(u *User) error {
	stmt := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(stmt, u.ID)
	return err
}

func Last() User {
	var u User
	stmt := "SELECT id first_name, last_name, email FROM users ORDER BY id desc LIMIT 1"
	row := db.QueryRow(stmt)
	row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	return u
}

func ResetStorage() {
	_, err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY")
	if err != nil {
		log.Fatal(err)
	}
}
