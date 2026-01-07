package middlewares

/* import (
	"log/slog"
	"net/http"
)

func FirstMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Hello world 1")
		next.ServeHTTP(w, r)
	}
}

func SecondMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Hello world 2")
		next.ServeHTTP(w, r)
	}
}

func ThirdMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Hello world 3")
		next.ServeHTTP(w, r)
	}
} */

// Abort middleware

/* func SecondMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Hello world 2")
		if true {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
} */
