package user

import (
	"database/sql"

	"net/http"
	"simple-project/utils"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get user data
	data := utils.GetJson(r)

	// search for this user
	var user User
	query := "SELECT username, password FROM users WHERE username = ?"
	err := db.QueryRow(query, data["username"].(string)).Scan(&user.Username, &user.Password)

	// check username or password
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			http.Error(w, "User Not Found", http.StatusBadRequest)
			return
		}
	}

	if !utils.CheckHashedPassword(data["password"].(string), user.Password) {
		http.Error(w, "Password is wrong", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Login Successfully"))

}

func SignUpHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// get user data
	data := utils.GetJson(r)

	// hashing password
	hashedPassword, err := utils.HashPassword(data["password"].(string))
	if err != nil {
		http.Error(w, "There is an error in password", http.StatusBadRequest)
		return
	}

	// create user to db
	_, Eerr := db.Exec("INSERT INTO users (first_name, last_name, username, password) values (?,?,?,?)", data["first_name"], data["last_name"], data["username"], hashedPassword)
	if Eerr != nil {
		if strings.Contains(Eerr.Error(), "UNIQUE constraint failed") {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
