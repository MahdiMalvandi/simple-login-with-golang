package user

import (
	"database/sql"
	"net/http"
	"path"
	"strings"
)

type UserRouter struct {
	DB     *sql.DB
	Prefix string
}

func (u *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := strings.TrimPrefix(path.Clean(r.URL.Path), u.Prefix)

	switch route {
	case "login":
		if r.Method == http.MethodPost {
			LoginHandler(w, r, u.DB)
			return
		}

	case "sign-up":
		if r.Method == http.MethodPost {
			SignUpHandler(w, r, u.DB)
			return
		}

	case "check-token":
		if r.Method == http.MethodPost {
			CheckJwtToken(w, r, u.DB)
			return
		}
	default:
		http.NotFound(w, r)
		return
	}

}
