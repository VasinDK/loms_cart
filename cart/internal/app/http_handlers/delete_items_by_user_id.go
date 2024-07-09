package http_handlers

import (
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/list/clear_cart"
	"time"
)

// DeleteItemsByUserID - удаляет все товары корзины по id пользователя
func (s *Server) DeleteItemsByUserID(h *clear_cart.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "DeleteItemsByUserID"
		const currentAddress = "DELETE /user/{user_id}/cart"
		var errExit = model.ErrOk
		var ctx = r.Context()

		ctx, span := tracer.Start(ctx, currentAddress)
		defer span.End()

		requestTotal.WithLabelValues(currentAddress).Inc()
		defer func(start time.Time) {
			requestTimeStatusUrl.WithLabelValues(errExit.Error(), currentAddress).Observe(time.Since(start).Seconds())
		}(time.Now())

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

		err = h.ClearCart(ctx, userId)
		if err != nil {
			logger.Errorw(ctx, op, "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			errExit = model.ErrHClearCart
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
