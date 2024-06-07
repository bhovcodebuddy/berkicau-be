package post

import (
	"github.com/gorilla/mux"
)

//init router
func PostRouter(router *mux.Router) *mux.Router {
	h := NewPostHandler()
	route := router.PathPrefix("/post").Subrouter()
	route.HandleFunc("", h.DoPost).Methods("POST")

	return router
}