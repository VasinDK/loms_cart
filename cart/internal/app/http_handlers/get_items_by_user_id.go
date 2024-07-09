package http_handlers

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/list/get_cart"
	"time"
)

// GetItemsByUserID - получает все товары корзины пользователя по id пользователя
func (s *Server) GetItemsByUserID(h *get_cart.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "GetItemsByUserID"
		const currentAddress = "GET /user/{user_id}/cart/list"
		var errExit = model.ErrOk
		var ctx = r.Context()

		ctx, span := tracer.Start(ctx, currentAddress)
		defer span.End()

		requestTotal.WithLabelValues(currentAddress).Inc()
		defer func(start time.Time) {
			requestTimeStatusUrl.WithLabelValues(errExit.Error(), currentAddress).Observe(time.Since(start).Seconds())
		}(time.Now())

		w.Header().Set("Content-Type", "application/json")

		userId, err := getPathValueInt(w, r, "user_id")
		if err != nil {
			errExit = model.ErrGetPathValueInt
			return
		}

		errs := validate.Var(userId, "required,gte=1")
		if errs != nil {
			logger.Errorw(ctx, op, "errs", errs)
			w.WriteHeader(http.StatusBadRequest)
			errExit = model.ErrValidateVar
			return
		}

		items, err := h.GetCart(ctx, userId)
		if err != nil || items.TotalPrice == 0 {
			if err != nil {
				logger.Errorw(ctx, op, "err", err)
				errExit = model.ErrHGetCart
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
			logger.Errorw(ctx, op, "err", err)
			w.WriteHeader(http.StatusNotFound)
			errExit = model.ErrJsonMarshal
			return
		}

		w.Write(buf)
	}
}
