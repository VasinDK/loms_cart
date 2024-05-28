package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"route256/cart/internal/pkg/cart/model"
	"route256/cart/internal/pkg/cart/service"
)

func (s *Server) AddItemCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getPathValueInt(w, r, "user_id")
	if err != nil {
		return
	}

	sku, err := getPathValueInt(w, r, "sku_id")
	if err != nil {
		return
	}

	productRequest := model.Product{}

	err = json.NewDecoder(r.Body).Decode(&productRequest)
	defer r.Body.Close()

	if err != nil || userId == 0 || sku == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productRequest.SKU = sku

	err = s.Service.AddProduct(&productRequest, userId)
	if errors.Is(err, service.ErrNoProductInStock) {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
