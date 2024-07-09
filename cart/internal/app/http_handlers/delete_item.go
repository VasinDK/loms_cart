package http_handlers

import (
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/item/delete_item"
	"time"
)

// DeleteItem - удаляет товар из корзины
func (s *Server) DeleteItem(h *delete_item.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "DeleteItem"
		const currentAddress = "DELETE /user/{user_id}/cart/{sku_id}"
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

		sku, err := getPathValueInt(w, r, "sku_id")
		if err != nil {
			errExit = model.ErrGetPathValueInt
			return
		}

		errs = validate.Var(sku, "required,gte=1")
		if errs != nil {
			logger.Errorw(ctx, op, "errs", errs)
			w.WriteHeader(http.StatusBadRequest)
			errExit = model.ErrValidateVar
			return
		}

		err = h.DeleteProductCart(ctx, userId, sku)
		if err != nil {
			logger.Errorw(ctx, op, "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			errExit = model.ErrHDeleteProductCart
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
