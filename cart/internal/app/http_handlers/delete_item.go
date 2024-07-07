package http_handlers

import (
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/item/delete_item"
	"time"

	"go.opentelemetry.io/otel"
)

// DeleteItem - удаляет товар из корзины
func (s *Server) DeleteItem(h *delete_item.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "DeleteItem"
		const currentAddress = "DELETE /user/{user_id}/cart/{sku_id}"
		var errExit = model.ErrOk
		var ctx = r.Context()

		tracer := otel.Tracer(model.ServiceName)
		ctx, span := tracer.Start(ctx, currentAddress)
		defer span.End()

		requestTotal.WithLabelValues(r.URL.Path).Inc()
		defer func(start time.Time) {
			requestTimeStatusUrl.WithLabelValues(errExit.Error(), r.URL.Path).Observe(time.Since(start).Seconds())
		}(time.Now())

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

		sku, err := getPathValueInt(w, r, "sku_id")
		if err != nil {
			errExit = err
			return
		}

		errs = validate.Var(sku, "required,gte=1")
		if errs != nil {
			logger.Errorw(ctx, op, "errs", errs)
			w.WriteHeader(http.StatusBadRequest)
			errExit = errs
			return
		}

		err = h.DeleteProductCart(ctx, userId, sku)
		if err != nil {
			logger.Errorw(ctx, op, "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			errExit = err
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
