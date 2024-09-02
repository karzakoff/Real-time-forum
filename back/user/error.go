package userFunctions

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	errorResponse := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(errorResponse)
}
