package main

import (
	"fmt"
	"net/http"
	"ngobar/berkicau/configs"
	"ngobar/berkicau/handlers/registration"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	//init db here
	_, err := configs.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()

	registration.RegistrationRouter(router)

	addr := "localhost:8080"
	fmt.Printf("Server listening on %s...\n", addr)
	http.ListenAndServe(addr, router)
}
