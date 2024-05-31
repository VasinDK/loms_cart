package middleware

import (
	"log/slog"
	"net/http"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if req := recover(); req != nil {
				w.WriteHeader(http.StatusInternalServerError)
				slog.Error("panic, PanicRecovery()")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
