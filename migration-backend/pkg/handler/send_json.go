package handler

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := encoder.Encode(body)
	if err != nil {
		panic(err)
	}
}
