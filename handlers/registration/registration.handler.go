package registration

import (
	"encoding/json"
	"net/http"
	"ngobar/berkicau/requests"
	"ngobar/berkicau/services"
)

type registrationHandler struct {
	userService services.UserService
}

func NewRegistrationHandler() registrationHandler {
	userService := services.NewUserService()
	return registrationHandler{userService: userService}
}

func generateResponse(writter http.ResponseWriter, httpCode int, resp interface{}) {
	response, _ := json.Marshal(resp)
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(httpCode)
	writter.Write(response)
}

func (h *registrationHandler) DoRegistration(writter http.ResponseWriter, request *http.Request) {

	var bodyRequest requests.RegistrationRequest

	//validate & decode body request
	err := json.NewDecoder(request.Body).Decode(&bodyRequest)
	if err != nil {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "Invalid request body",
		}
		generateResponse(writter, http.StatusBadRequest, responseData)
		return
	}

	status, msg := h.userService.RegisterNewUser(bodyRequest)

	responseData := map[string]interface{}{
		"status":  status,
		"message": msg,
	}

	generateResponse(writter, http.StatusOK, responseData)
}
