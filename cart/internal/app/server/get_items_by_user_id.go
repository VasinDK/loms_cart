package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Получает все товары корзины пользователя по id пользователя
func (s *Server) GetItemsByUserID(w http.ResponseWriter, r *http.Request) {
	op := "GetItemsByUserID"

	w.Header().Set("Content-Type", "application/json")

	userId, err := getPathValueInt(w, r, "user_id")
	if err != nil {
		return
	}

	items, err := s.Service.GetCart(userId)
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
