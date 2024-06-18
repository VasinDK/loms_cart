package http_handlers

import (
	"context"
	"log/slog"
	"net/http"
	"route256/cart/internal/service/list/clear_cart"
)

// DeleteItemsByUserID - удаляет все товары корзины по id пользователя
func (s *Server) DeleteItemsByUserID(h *clear_cart.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "DeleteItemsByUserID"

		userId, err := getPathValueInt(w, r, "user_id")
		if err != nil {
			return
		}

		errs := validate.Var(userId, "required,gte=1")
		if errs != nil {
			slog.Error(op, errs)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.ClearCart(context.Background(), userId)
		if err != nil {
			slog.Error(op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
