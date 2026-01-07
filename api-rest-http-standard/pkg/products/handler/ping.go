package v2

import (
	"example.com/shared/api/middlewares"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// Ping godoc
// @Summary Health check
// @Description Returns pong if server is alive
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string "message: pong"
// @Failure 500 {object} map[string]string "error: user id not found"
// @Router /ping [get]
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