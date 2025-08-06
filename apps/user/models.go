package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT 0
);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("db ", err)
	}
}

func (u *User) CreateUser(db *sql.DB) error {
	// create user to db
	_, Eerr := db.Exec("INSERT INTO users (first_name, last_name, username, password) values (?,?,?,?)", u.FirstName, u.LastName, u.Username, u.Password)
	if Eerr != nil {
		if strings.Contains(Eerr.Error(), "UNIQUE constraint failed") {
			return errors.New("User Already Exists")
		}
	}
	return nil
}

func (u *User) Search(db *sql.DB, username string) error {
	// create user to db
	query := "SELECT id, username, password FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&u.Id, &u.Username, &u.Password)

	return err
}

func (u *User) GetAll(db *sql.DB) ([]User, error) {
	query := "SELECT id, first_name, last_name, username, is_admin FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		if Rerr := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.IsAdmin); Rerr != nil {
			return users, Rerr
		}
		users = append(users, user)
	}

	return users, nil
}
func (u *User) GetUserByField(db *sql.DB, field string, data string) (User, error) {
	allowedFields := map[string]bool{
		"username":   true,
		"id":         true,
		"first_name": true,
		"last_name":  true,
	}

	if !allowedFields[field] {
		return User{}, errors.New("field Must be one of the User fields")
	}

	query := fmt.Sprintf("SELECT id, first_name, last_name, username, is_admin from users WHERE %s = ?", field)
	row := db.QueryRow(query, data)

	err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username, &u.IsAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}
	return *u, err
}
