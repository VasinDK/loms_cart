package http_handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/service/item/add_product"
)

// AddItem - добавляет товар в корзину
func (s *Server) AddItem(h *add_product.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "AddItem"

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

		productRequest := ProductRequest{}

		err = json.NewDecoder(r.Body).Decode(&productRequest)
		if err != nil {
			slog.Error(op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		errs = validate.Struct(productRequest)
		if errs != nil {
			slog.Error(op, errs)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product := model.Product{}

		product.SKU = sku
		product.Count = productRequest.Count

		err = h.AddProduct(context.Background(), &product, userId)

		if errors.Is(err, model.ErrNoProductInStock) {
			w.WriteHeader(http.StatusPreconditionFailed)
			return
		}

		if errors.Is(err, model.ErrInsufficientStock) {
			w.WriteHeader(http.StatusPreconditionFailed)
			w.Write([]byte(model.ErrInsufficientStock.Error()))
			return
		}

		if err != nil {
			slog.Error(op, err)
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
