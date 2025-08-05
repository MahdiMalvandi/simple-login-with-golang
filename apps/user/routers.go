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

	case "sign-up":
		if r.Method == http.MethodPost {
			SignUpHandler(w, r, u.DB, u.Logger)
			return
		}

	case "check-token":
		if r.Method == http.MethodPost {
			CheckJwtToken(w, r, u.DB, u.Logger)
			return
		}
	default:
		http.NotFound(w, r)
		return
	}

}
