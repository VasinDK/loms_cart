package http_handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"route256/cart/internal/service/list/get_cart"
)

// Получает все товары корзины пользователя по id пользователя
func (s *Server) GetItemsByUserID(h *get_cart.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "GetItemsByUserID"

		w.Header().Set("Content-Type", "application/json")

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

		items, err := h.GetCart(userId)
		if err != nil || items.TotalPrice == 0 {
			if err != nil {
				slog.Error(op, err)
			}
			w.WriteHeader(http.StatusNotFound)
			return
		}

		cartResponse := CartResponse{}

		for i := range items.Items {
			ProductRes := ProductResponse{
				Count: items.Items[i].Count,
				Price: items.Items[i].Price,
				Name:  items.Items[i].Name,
				SKU:   items.Items[i].SKU,
			}
			cartResponse.Items = append(cartResponse.Items, &ProductRes)
		}

		cartResponse.TotalPrice = items.TotalPrice

		buf, err := json.Marshal(cartResponse)
		if err != nil {
			slog.Error(op, err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write(buf)
	}
}
