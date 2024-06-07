package repost

import (
	"github.com/gorilla/mux"
)

//init router
func RepostRouter(router *mux.Router) *mux.Router {
	h := NewRepostHandler()
	route := router.PathPrefix("/repost").Subrouter()
	route.HandleFunc("", h.DoRepost).Methods("POST")

	return router
}