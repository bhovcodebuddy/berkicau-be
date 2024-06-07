package login

import (
	"net/http"
	"ngobar/berkicau/services"
)

type loginHandler struct {
	userService services.UserService
}

func NewLoginHandler() loginHandler {
	return loginHandler{userService: services.NewUserService()}
}

func (h *loginHandler) DoLogin(writter http.ResponseWriter, request *http.Request) {

}