package server

import (
	"log/slog"
	"net/http"
)

// Удаляет товар из корзины
func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {
	op := "DeleteItem"

	userId, err := getPathValueInt(w, r, "user_id")
	if err != nil {
		return
	}

	sku, err := getPathValueInt(w, r, "sku_id")
	if err != nil {
		return
	}

	err = s.Service.DeleteProductCart(userId, sku)
	if err != nil {
		slog.Error(op, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
