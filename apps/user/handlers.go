package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple-project/utils"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get user data
	data := utils.GetJson(r)

	// search for this user
	query := "SELECT username, password FROM users WHERE username = ?"
	row := db.QueryRow(query, data["username"])

	fmt.Print(row)
	// check username or password
	// if it was true -> create jwt token
	// if it wast true -> error

}

func SignUpHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	
	data := utils.GetJson(r)
	hashedPassword, err := utils.HashPassword(data["password"].(string))
	if err != nil {
		http.Error(w, "There is an error in password", http.StatusBadRequest)
			return 

	}
	_, Eerr := db.Exec("INSERT INTO users (first_name, last_name, username, password) values (?,?,?,?)", data["first_name"], data["last_name"], data["username"], hashedPassword)
	if Eerr != nil {
		if strings.Contains(Eerr.Error(), "UNIQUE constraint failed"){
			http.Error(w, "User already exists", http.StatusBadRequest)
			return 
		}
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
