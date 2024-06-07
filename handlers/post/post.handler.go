package post

import (
	"net/http"
	"ngobar/berkicau/services"
)

type postHandler struct {
	postService services.PostService
}

func NewPostHandler() postHandler {
	return postHandler{postService: services.NewPostService()}
}

func (h *postHandler) DoPost(writter http.ResponseWriter, request *http.Request) {

}