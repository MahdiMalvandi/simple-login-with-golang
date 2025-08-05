package api

import (
	"net/http"
	"simple-project/apps/user"
	"simple-project/configs"
	"simple-project/middlewares"
	"simple-project/utils"
)

func SetupRoutes(mux *http.ServeMux, logger *utils.Logger){
	middleware := middlewares.Middleware{}
	mux.Handle("/auth/", middleware.JsonParser(&user.UserRouter{DB: configs.DB, Prefix:"/auth/", Logger: logger}, logger))
	
}