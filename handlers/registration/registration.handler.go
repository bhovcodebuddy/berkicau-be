package registration

import (
	"encoding/json"
	"net/http"
)

type RegistrationHandler struct {
}

func DoRegistration(writter http.ResponseWriter, request *http.Request) {

	//todo: insert here

	response := map[string]interface{}{
		"status":  true,
		"message": "Data found.",
		"data":    nil,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		// Handle error
	}

	// Set content type and status code
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusOK)

	// Write JSON response
	writter.Write(jsonData)
}
