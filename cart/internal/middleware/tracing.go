package middleware

import (
	"net/http"

	"go.opentelemetry.io/otel"
)

const currentServiceName = "Cart"
const ignorePath = "/metrics"

func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != ignorePath {
			tracer := otel.Tracer(currentServiceName)
			ctx, span := tracer.Start(r.Context(), currentServiceName)
			defer span.End()

			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
