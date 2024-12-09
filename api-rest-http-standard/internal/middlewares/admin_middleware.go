package middlewares

import (
	"log/slog"
	"net/http"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Start", "Method", r.Method, "Url", r.URL.Path)
		defer slog.Info("End", "Method", r.Method, "Url", r.URL.Path)
		next.ServeHTTP(w,r)
	}
}




func AuthRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Pase por aqui......................")
		next(w,r)
	}
}
