package http_handlers

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/list/checkout"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type User struct {
	User int64 `json:"user"`
}

type Order struct {
	OrderId int64 `json:"orderId"`
}

// Checkout - создает ордер
func (s *Server) Checkout(h *checkout.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "Checkout"
		const currentAddress = "POST /user/cart/checkout"
		var errExit = model.ErrOk
		var ctx = r.Context()

		ctx, span := tracer.Start(ctx, currentAddress)
		defer span.End()

		requestTotal.WithLabelValues(currentAddress).Inc()
		defer func(start time.Time) {
			requestTimeStatusUrl.WithLabelValues(errExit.Error(), currentAddress).Observe(time.Since(start).Seconds())
		}(time.Now())

		w.Header().Set("Content-Type", "application/json")

		user := &User{}

		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			logger.Errorw(ctx, op, "Checkout", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			errExit = model.ErrJsonNewDecoder
			return
		}

		orderId, err := h.Checkout(ctx, user.User)
		if err != nil {
			logger.Errorw(ctx, op, "h.Checkout", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			errExit = model.ErrHCheckout
			return
		}

		span.SetAttributes(attribute.Int64("orderId", orderId))

		order := &Order{orderId}
		buf, err := json.Marshal(order)
		if err != nil {
			logger.Errorw(ctx, op, "h.Checkout", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			errExit = model.ErrJsonMarshal
			return
		}

		w.Write(buf)
	}
}
