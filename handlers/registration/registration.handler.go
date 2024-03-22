package registration

import (
	"encoding/json"
	"net/http"
	"ngobar/berkicau/requests"
)

type RegistrationHandler struct {
}

func GenerateResponse(writter http.ResponseWriter, httpCode int, resp interface{}) {
	response, _ := json.Marshal(resp)
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(httpCode)
	writter.Write(response)
}

func DoRegistration(writter http.ResponseWriter, request *http.Request) {

	var bodyRequest requests.RegistrationRequest

	//validate & decode body request
	err := json.NewDecoder(request.Body).Decode(&bodyRequest)
	if err != nil {
		responseData := map[string]interface{}{
			"status":  false,
			"message": "Invalid request body",
		}
		GenerateResponse(writter, http.StatusBadRequest, responseData)
		return
	}

	//do register here

	responseData := map[string]interface{}{
		"status":  true,
		"message": "Registration succeed",
	}

	GenerateResponse(writter, http.StatusOK, responseData)
}
