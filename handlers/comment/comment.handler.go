package comment

import (
	"net/http"
	"ngobar/berkicau/services"
)

type commentHandler struct {
	postService services.PostService
}

func NewCommentHandler() commentHandler {
	return commentHandler{postService: services.NewPostService()}
}

func (h *commentHandler) DoComment(writter http.ResponseWriter, request *http.Request) {

}