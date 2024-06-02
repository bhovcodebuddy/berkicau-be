package registration

import (
	"github.com/gorilla/mux"
)

//init router
func RegistrationRouter(router *mux.Router) *mux.Router {
	h := NewRegistrationHandler()
	route := router.PathPrefix("/register").Subrouter()
	route.HandleFunc("", h.DoRegistration).Methods("POST")

	return router
}
