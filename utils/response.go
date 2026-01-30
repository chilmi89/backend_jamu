package utils

import (
	"encoding/json"
	"net/http"
)

// Response standar untuk API
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := APIResponse{
		Success: status >= 200 && status < 300,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func JSONError(w http.ResponseWriter, status int, message string) {
	JSONResponse(w, status, message, nil)
}
