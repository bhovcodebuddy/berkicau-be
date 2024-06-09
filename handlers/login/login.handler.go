package login

import (
	"encoding/json"
	"net/http"
	"ngobar/berkicau/requests"
	"ngobar/berkicau/services"
	"ngobar/berkicau/utils/http"
)

type loginHandler struct {
	userService services.UserService
}

func NewLoginHandler() loginHandler {
	return loginHandler{userService: services.NewUserService()}
}

func (h *loginHandler) DoLogin(writter http.ResponseWriter, request *http.Request) {
	var bodyRequest requests.LoginRequest

	//validate & decode body request
	err := json.NewDecoder(request.Body).Decode(&bodyRequest)
	if err != nil {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "Invalid request body",
		}
		utils.GenerateResponse(writter, http.StatusBadRequest, responseData)
		return
	}

	status, msg, token := h.userService.GetLoginInformation(bodyRequest)

	responseData := map[string]interface{}{
		"status":  status,
		"message": msg,
		"token": &token,
	}

	utils.GenerateResponse(writter, http.StatusOK, responseData)
}