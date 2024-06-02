package main

import (
	"fmt"
	"net/http"
	"ngobar/berkicau/configs"
	"ngobar/berkicau/handlers/registration"
	"ngobar/berkicau/helper"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	//init db here
	db, err := configs.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	helper.DB = db
	
	//init routers and handlers
	router := mux.NewRouter()
	registration.RegistrationRouter(router)

	addr := "localhost:8080"
	fmt.Printf("Server listening on %s...\n", addr)
	http.ListenAndServe(addr, router)
}
