package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"example.com/shared/api/middlewares"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	slog.Info("Ping handler: executing handler")

	reqID, ok := r.Context().Value(middlewares.RequestIDKey).(string)
	if !ok || reqID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Info("user id not found")
		return
	}
	slog.Info(fmt.Sprintf("Ping handler, user id: %s", reqID))

	w.WriteHeader(http.StatusOK)
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "pong",
	}
	json.NewEncoder(w).Encode((resp))
}
