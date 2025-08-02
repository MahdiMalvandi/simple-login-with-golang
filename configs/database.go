package configs

import (
	"database/sql"
	"log"
	"simple-project/apps/user"

    _ "modernc.org/sqlite"
)
var DB *sql.DB

func ServeDatabase() *sql.DB {
	var err error
	DB, err = sql.Open("sqlite", "file:database.db")
	if err != nil {
		log.Fatal(err)
	}
	return DB

}

func CreateTables(db *sql.DB) {
	// Creating tables
	user.CreateTable(db)
}
