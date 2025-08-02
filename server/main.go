package main

import (
	"net/http"
	"simple-project/configs"
	"simple-project/api"
	
)

func main() {


	// Connecting to Database
	db := configs.ServeDatabase()
	configs.CreateTables(db)


	// Making a server
	mux := http.NewServeMux()

	api.SetupRoutes(mux)

	err := http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		panic(err)
	}
}