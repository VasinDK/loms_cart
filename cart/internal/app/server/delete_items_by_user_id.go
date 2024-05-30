package server

import (
	"log/slog"
	"net/http"
)

// Удаляет все товары корзины по id пользователя
func (s *Server) DeleteItemsByUserID(w http.ResponseWriter, r *http.Request) {
	op := "DeleteItemsByUserID"

	userId, err := getPathValueInt(w, r, "user_id")
	if err != nil {
		return
	}

	err = s.Service.ClearCart(userId)
	if err != nil {
		slog.Error(op, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
