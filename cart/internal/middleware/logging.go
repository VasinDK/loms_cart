package middleware

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

// Logging - middleware логирующий запрос, время ответа
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PRI" && r.RequestURI == "*" {
			// Игнорируем запросы с методом PRI и URI *
			return
		}

		start := time.Now()

		head := []string{}
		for name, values := range r.Header {
			head = append(head, fmt.Sprintf("%v ", name))
			for _, value := range values {
				head = append(head, fmt.Sprintf("%v ", value))
			}
		}

		body := []byte("")
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			err := fmt.Errorf("")
			body, err = io.ReadAll(r.Body)

			if err != nil {
				slog.Info("Request",
					slog.String("URI", r.RequestURI),
					slog.String("Method", r.Method),
					slog.String("Header", strings.Join(head, "")),
					slog.String("err", err.Error()),
				)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		r.Body = io.NopCloser(io.Reader(strings.NewReader(string(body))))

		next.ServeHTTP(w, r)

		slog.Info("Request",
			slog.String("URI", r.RequestURI),
			slog.String("Method", r.Method),
			slog.String("Body", string(body)),
			slog.String("Header", strings.Join(head, "")),
			slog.Duration("Время обработки запроса", time.Since(start)),
		)
	})
}
