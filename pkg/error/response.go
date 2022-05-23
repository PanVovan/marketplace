package error_response

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse{
		Message: message,
	})
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewStatusResponse(status string) statusResponse {
	return statusResponse{Status: status}
}
