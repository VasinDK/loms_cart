package http_handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/service/item/add_product"
)

// Добавляет товар в корзину.
func (s *Server) AddItem(h *add_product.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "AddItem"

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
		defer r.Body.Close()

		errs = validate.Struct(productRequest)
		if errs != nil {
			slog.Error(op, errs)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product := model.Product{}

		product.SKU = sku
		product.Count = productRequest.Count

		err = h.AddProduct(&product, userId)
		if errors.Is(err, model.ErrNoProductInStock) {
			w.WriteHeader(http.StatusPreconditionFailed)
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
