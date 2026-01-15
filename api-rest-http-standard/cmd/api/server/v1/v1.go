package v1

import (
	"encoding/json"
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// Api version
	v1 := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	// Here's everything your app needs
	// database connection 

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := struct {
			Msg string `json:"msg"`
		}{
			Msg: "pong",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})
	return mux
}
