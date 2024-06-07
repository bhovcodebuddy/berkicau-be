package utils

import (
	"encoding/json"
	"net/http"
)

func GenerateResponse(writter http.ResponseWriter, httpCode int, resp interface{}) {
	response, _ := json.Marshal(resp)
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(httpCode)
	writter.Write(response)
}