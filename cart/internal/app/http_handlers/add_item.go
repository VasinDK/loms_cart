package http_handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/service/item/add_product"
	"time"

	"go.opentelemetry.io/otel"
)

// AddItem - добавляет товар в корзину
func (s *Server) AddItem(h *add_product.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "AddItem"
		const currentAddress = "POST /user/{user_id}/cart/{sku_id}"
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

		productRequest := ProductRequest{}

		err = json.NewDecoder(r.Body).Decode(&productRequest)
		if err != nil {
			logger.Errorw(ctx, op, "errs", errs)
			w.WriteHeader(http.StatusBadRequest)
			errExit = err
			return
		}

		errs = validate.Struct(productRequest)
		if errs != nil {
			logger.Errorw(ctx, op, "errs", errs)
			w.WriteHeader(http.StatusBadRequest)
			errExit = errs
			return
		}

		product := model.Product{}

		product.SKU = sku
		product.Count = productRequest.Count

		err = h.AddProduct(ctx, &product, userId)
		if err != nil {
			errExit = err
		}

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
			logger.Errorw(ctx, op, "errs", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
