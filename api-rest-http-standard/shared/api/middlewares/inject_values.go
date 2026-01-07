package middlewares

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

type RequestKey string

var RequestIDKey RequestKey = "request-id"

func FirstMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := "user_id"
		ctx := context.WithValue(r.Context(), RequestIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func SecondMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID, ok := r.Context().Value(RequestIDKey).(string)
		if !ok || reqID == "" {
            w.WriteHeader(http.StatusInternalServerError)
			slog.Info("user id not found")
			return
		}
		slog.Info(fmt.Sprintf("SecondMdw, user id: %s", reqID))
		next.ServeHTTP(w, r)
	}
}

func ThirdMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Hello world 3")
		next.ServeHTTP(w, r)
	}
}
