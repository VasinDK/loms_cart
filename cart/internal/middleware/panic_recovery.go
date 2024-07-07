package middleware

import (
	"context"
	"net/http"
	"route256/cart/internal/pkg/logger"
)

// PanicRecovery - middleware восстанавливающий работу приложения в случае появления паники
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if req := recover(); req != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Errorw(context.Background(), "panic, PanicRecovery()")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
