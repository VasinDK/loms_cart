package grpc_handlers

import (
	"context"
	"errors"
	"route256/loms/internal/model"
	"route256/loms/internal/pkg/logger"
	"route256/loms/pkg/api/loms/v1"
	"route256/loms/pkg/statuses"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OrderPay - покупка
func (h *Handlers) OrderPay(ctx context.Context, orderId *loms.OrderPayRequest) (*loms.OrderPayResponse, error) {
	const op = "OrderPay"
	var errExit error

	ctx, span := tracer.Start(ctx, op)
	defer span.End()

	requestTotal.WithLabelValues(op).Inc()

	defer func(start time.Time) {
		requestTimeStatusUrl.WithLabelValues(
			statuses.GetStatusGRPC(errExit),
			op,
		).Observe(time.Since(start).Seconds())
	}(time.Now())

	err := h.service.OrderPay(ctx, model.OrderId(orderId.GetOrderId()))

	if errors.Is(err, model.ErrStatusNoAwaitingPayment) {
		logger.Errorw(ctx, op, "h.service.OrderPay", "err", err)
		errExit = status.Error(codes.NotFound, model.ErrStatusNoAwaitingPayment.Error())
		return &loms.OrderPayResponse{}, errExit
	}

	if errors.Is(err, model.ErrSkuNoSuch) {
		logger.Errorw(ctx, op, "h.service.OrderPay", "err", err)
		errExit = status.Error(codes.InvalidArgument, model.ErrSkuNoSuch.Error())
		return &loms.OrderPayResponse{}, errExit
	}

	if err != nil {
		logger.Errorw(ctx, op, "h.service.OrderPay", "err", err)
		errExit = status.Error(codes.Internal, model.ErrSkuNoSuch.Error())
		return &loms.OrderPayResponse{}, errExit
	}

	return &loms.OrderPayResponse{}, errExit
}
