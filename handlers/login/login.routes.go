package login

import (
	"github.com/gorilla/mux"
)

//init router
func LoginRouter(router *mux.Router) *mux.Router {
	h := NewLoginHandler()
	route := router.PathPrefix("/login").Subrouter()
	route.HandleFunc("", h.DoLogin).Methods("POST")

	return router
}
