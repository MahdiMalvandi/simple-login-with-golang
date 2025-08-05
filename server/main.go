package main

import (
	"net/http"
	"simple-project/configs"
	"simple-project/api"
	"simple-project/utils"
	
)

func main() {


	// Connecting to Database
	db := configs.ServeDatabase()
	configs.CreateTables(db)

	// Creating Logger 
	logger := utils.NewLogger()

	// Making a server
	mux := http.NewServeMux()

	api.SetupRoutes(mux, logger)

	err := http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		panic(err)
	}
}