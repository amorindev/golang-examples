package handler

import (
	"encoding/json"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "pong",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((resp))
}
