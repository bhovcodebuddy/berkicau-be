package repost

import (
	"net/http"
	"ngobar/berkicau/services"
)

type repostHandler struct {
	postService services.PostService
}

func NewRepostHandler() repostHandler {
	return repostHandler{postService: services.NewPostService()}
}

func (h *repostHandler) DoRepost(writter http.ResponseWriter, request *http.Request) {

}