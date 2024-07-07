package http_handlers

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/list/get_cart"
	"time"

	"go.opentelemetry.io/otel"
)

// GetItemsByUserID - получает все товары корзины пользователя по id пользователя
func (s *Server) GetItemsByUserID(h *get_cart.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "GetItemsByUserID"
		const currentAddress = "GET /user/{user_id}/cart/list"
		var errExit = model.ErrOk
		var ctx = r.Context()

		tracer := otel.Tracer(model.ServiceName)
		ctx, span := tracer.Start(ctx, currentAddress)
		defer span.End()

		requestTotal.WithLabelValues(r.URL.Path).Inc()
		defer func(start time.Time) {
			requestTimeStatusUrl.WithLabelValues(errExit.Error(), r.URL.Path).Observe(time.Since(start).Seconds())
		}(time.Now())

		w.Header().Set("Content-Type", "application/json")

		userId, err := getPathValueInt(w, r, "user_id")
		if err != nil {
			errExit = err
			return
		}

		errs := validate.Var(userId, "required,gte=1")
		if errs != nil {
			logger.Errorw(ctx, op, "errs", errs)
			w.WriteHeader(http.StatusBadRequest)
			errExit = errs
			return
		}

		items, err := h.GetCart(ctx, userId)
		if err != nil || items.TotalPrice == 0 {
			if err != nil {
				logger.Errorw(ctx, op, "err", err)
				errExit = err
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
			errExit = err
			return
		}

		w.Write(buf)
	}
}
