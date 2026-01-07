package middlewares

import (
	"log/slog"
	"net/http"
)

func MyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("MyMiddleware: preparing the execution of the final handler")

		next.ServeHTTP(w, r)

		slog.Info("MyMiddleware: the handler has finished")
	})
}

func MyHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("MyHandlerMiddleware: preparing the execution of the final handler")

		next.ServeHTTP(w, r)

		slog.Info("MyHandlerMiddleware: the handler has finished")
	})
}


