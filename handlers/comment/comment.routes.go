package comment

import (
	"github.com/gorilla/mux"
)

//init router
func CommentRouter(router *mux.Router) *mux.Router {
	h := NewCommentHandler()
	route := router.PathPrefix("/comment").Subrouter()
	route.HandleFunc("", h.DoComment).Methods("POST")

	return router
}