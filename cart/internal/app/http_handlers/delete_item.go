package http_handlers

import (
	"log/slog"
	"net/http"
	"route256/cart/internal/service/item/delete_item"
)

// DeleteItem - удаляет товар из корзины
func (s *Server) DeleteItem(h *delete_item.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "DeleteItem"
		ctx := r.Context()

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

		sku, err := getPathValueInt(w, r, "sku_id")
		if err != nil {
			return
		}

		errs = validate.Var(sku, "required,gte=1")
		if errs != nil {
			slog.Error(op, errs)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.DeleteProductCart(ctx, userId, sku)
		if err != nil {
			slog.Error(op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
