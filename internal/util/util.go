package util

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, payload interface{}, statusCode int) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}
