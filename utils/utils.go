package utils

import (
	"encoding/json"
	"net/http"
)

// Exception : struct for json error messages to be returned
type Exception struct {
	Message string `json:"message"`
}

// WriteJSONMessage : write message with json header
func WriteJSONMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Exception{Message: message})
}

