package user

import (
	"database/sql"
	"fmt"
	"net/http"
)
type UserRouter struct{
	DB *sql.DB
}

func (u *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/login":
			if r.Method == http.MethodPost{
				LoginHandler(w , r, u.DB)
				return
			}
			

		case "/sign-up":
			if r.Method == http.MethodPost{
				SignUpHandler(w , r, u.DB)
				return
			}

		
	}
	fmt.Println("error")
	return

}
