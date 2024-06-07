package like

import (
	"net/http"
)

type likeHandler struct {
}

func NewLikeHandler() likeHandler {
	return likeHandler{}
}

func (h *likeHandler) DoLike(writter http.ResponseWriter, request *http.Request) {

}