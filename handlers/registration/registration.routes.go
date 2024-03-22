package registration

import "github.com/gorilla/mux"

//init router
func RegistrationRouter(router *mux.Router) *mux.Router {

	route := router.PathPrefix("/register").Subrouter()
	route.HandleFunc("", DoRegistration).Methods("POST")

	return route
}
