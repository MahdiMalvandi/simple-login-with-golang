package api

import (
	"net/http"
	"simple-project/apps/user"
	"simple-project/configs"
	"simple-project/middlewares"
)

func SetupRoutes(mux *http.ServeMux){
	middleware := middlewares.Middleware{}
	mux.Handle("/auth/", middleware.JsonParser(&user.UserRouter{DB: configs.DB, Prefix:"/auth/"}))
	
}