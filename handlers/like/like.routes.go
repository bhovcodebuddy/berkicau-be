package like

import (
	"github.com/gorilla/mux"
)

//init router
func LikeRouter(router *mux.Router) *mux.Router {
	h := NewLikeHandler()
	route := router.PathPrefix("/like").Subrouter()
	route.HandleFunc("", h.DoLike).Methods("POST")

	return router
}