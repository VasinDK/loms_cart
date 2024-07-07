package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"route256/cart/internal/pkg/logger"
	"strings"
	"time"
)

// Logging - middleware логирующий запрос, время ответа
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

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
				logger.Infow(
					ctx,
					"Request",
					"URI", r.RequestURI,
					"Method", r.Method,
					"Header", strings.Join(head, ""),
					"err", err.Error(),
				)

				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		r.Body = io.NopCloser(io.Reader(strings.NewReader(string(body))))

		next.ServeHTTP(w, r)

		logger.Infow(
			ctx,
			"Request",
			"URI", r.RequestURI,
			"Method", r.Method,
			"Body", string(body),
			"Header", strings.Join(head, ""),
			"Время обработки запроса", time.Since(start),
		)
	})
}
