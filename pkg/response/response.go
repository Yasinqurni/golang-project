package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, message string, data any, meta any, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	response := ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	writeJSONResponse(w, statusCode, response)
}

func ErrorResponse(w http.ResponseWriter, message string, errors any, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}
	response := ApiResponse{
		Status:  "error",
		Message: message,
		Errors:  errors,
	}

	if statusCode == http.StatusInternalServerError {
		response.Errors = nil
	}

	writeJSONResponse(w, statusCode, response)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, response ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
