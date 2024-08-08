package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) (int, error) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return status, json.NewEncoder(w).Encode(v)
}

func WriteErrorJSON(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	WriteJSON(w, code, errResponse{Error: msg})
}
