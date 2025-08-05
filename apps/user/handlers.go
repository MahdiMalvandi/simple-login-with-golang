package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-project/utils"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, logger *utils.Logger) {
	// get user data
	data := utils.GetJson(r)

	// search for this user
	var user User
	query := "SELECT id, username, password FROM users WHERE username = ?"
	err := db.QueryRow(query, data["username"].(string)).Scan(&user.Id, &user.Username, &user.Password)

	// check username or password
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			http.Error(w, "User Not Found", http.StatusBadRequest)
			logger.Log(fmt.Sprintf("ERROR:User-Handlers:LoginHandler->User Does not Exists for username: '%s'", data["username"]))
			return
		}
	}

	// check user's password
	if !utils.CheckHashedPassword(data["password"].(string), user.Password) {
		http.Error(w, "Password is wrong", http.StatusBadRequest)
		logger.Log(fmt.Sprintf("ERROR:User-Handlers:LoginHandler->User's Password is wrong for username: '%s'", data["username"]))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	logger.Log(fmt.Sprintf("INFO:User-Handlers:LoginHandler->User Login was successful for username: '%s'", data["username"]))

	// creating jwt for user
	res, err := utils.CreateJwt(user.Id, user.Username)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR:User-Handlers:LoginHandler->Error in create jwt for username: '%s'", data["username"]))
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	logger.Log(fmt.Sprintf("INFO:User-Handlers:LoginHandler->User's Jwt was created for username: '%s'", data["username"]))
	w.Write([]byte(res))
}

func SignUpHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, logger *utils.Logger) {
	// get user data
	data := utils.GetJson(r)

	// hashing password
	hashedPassword, err := utils.HashPassword(data["password"].(string))
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR:User-Handlers:SignUpHandler->There is an error in hashing password for username: '%s'", data["username"]))
		http.Error(w, "There is an error in password", http.StatusBadRequest)
		return
	}

	// create user to db
	_, Eerr := db.Exec("INSERT INTO users (first_name, last_name, username, password) values (?,?,?,?)", data["first_name"], data["last_name"], data["username"], hashedPassword)
	if Eerr != nil {
		if strings.Contains(Eerr.Error(), "UNIQUE constraint failed") {
			logger.Log(fmt.Sprintf("ERROR:User-Handlers:SignUpHandler->User already exists for username: '%s'", data["username"]))

			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}
	}
	logger.Log(fmt.Sprintf("INFO:User-Handlers:SignUpHandler->User created successfully for username: '%s'", data["username"]))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func CheckJwtToken(w http.ResponseWriter, r *http.Request, db *sql.DB, logger *utils.Logger) {
	data := utils.GetJson(r)
	res, err := utils.VerifyJwt(data["token"].(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Log(fmt.Sprintf("ERROR:User-Handlers:CheckJwtToken->%s for username: '%s'", err.Error(), data["username"]))

		return
	}
	w.WriteHeader(http.StatusAccepted)
	logger.Log(fmt.Sprintf("INFO:User-Handlers:CheckJwtToken->Jwt checked successfully for username: '%s'", data["username"]))

	jsonData, _ := json.Marshal(map[string]bool{"status": res})
	w.Write([]byte(jsonData))
}
