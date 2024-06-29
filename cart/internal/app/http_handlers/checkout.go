package http_handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"route256/cart/internal/service/list/checkout"
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

		ctx := r.Context()

		w.Header().Set("Content-Type", "application/json")

		user := &User{}

		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			slog.Error(op, "Checkout", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		orderId, err := h.Checkout(ctx, user.User)
		if err != nil {
			slog.Error(op, "h.Checkout", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		order := &Order{orderId}
		buf, err := json.Marshal(order)
		if err != nil {
			slog.Error(op, "h.Checkout", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write(buf)
	}
}
