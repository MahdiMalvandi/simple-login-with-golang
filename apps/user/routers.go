package user

import (
	"database/sql"
	"net/http"
	"path"
	"simple-project/utils"
	"strings"
)

type UserRouter struct {
	DB     *sql.DB
	Prefix string
	Logger *utils.Logger
}

func (u *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := strings.TrimPrefix(path.Clean(r.URL.Path), u.Prefix)

	switch route {
	case "login":
		if r.Method == http.MethodPost {
			LoginHandler(w, r, u.DB, u.Logger)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	case "sign-up":
		if r.Method == http.MethodPost {
			SignUpHandler(w, r, u.DB, u.Logger)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	case "check-token":
		if r.Method == http.MethodPost {
			CheckJwtToken(w, r, u.DB, u.Logger)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	case "users":
		if r.Method == http.MethodGet {
			GetListOfUsers(w, r, u.DB, u.Logger)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	case "get-user":
				if r.Method == http.MethodGet {
			GetUserByField(w, r, u.DB, u.Logger)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	default:

		http.NotFound(w, r)
		return
	}

}
