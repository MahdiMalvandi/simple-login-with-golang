package user

import (
	"database/sql"
	"log"

    _ "modernc.org/sqlite"
)


type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	IsAdmin bool `json:"is_admin"`
}

func CreateTable(db *sql.DB){
	query := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT 0
);`
	_ , err := db.Exec(query)
	if err != nil {
		log.Fatal("db ",err)
	}
}

