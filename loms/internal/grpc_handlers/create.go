package grpc_handlers

import (
	"context"
	"route256/loms/internal/pkg/logger"
	"route256/loms/pkg/api/loms/v1"
	"route256/loms/pkg/statuses"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OrderCreate - создает ордер
func (h *Handlers) OrderCreate(ctx context.Context, order *loms.OrderCreateRequest) (*loms.OrderCreateResponse, error) {
	const op = "OrderCreate"
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

	orderModel, err := h.RepackOrderToModel(order)
	if err != nil {
		logger.Errorw(ctx, op, "h.RepackOrderInOrderModel", "err", err)
		errExit = status.Error(codes.FailedPrecondition, "so")
		return nil, errExit
	}

	orderIdModel, err := h.service.Create(ctx, orderModel)
	if err != nil {
		logger.Errorw(ctx, op, "h.service.OrderCreate", "err", err)
		errExit = status.Error(codes.FailedPrecondition, "I am sorry")
		return nil, errExit
	}

	return &loms.OrderCreateResponse{
		OrderId: int64(orderIdModel),
	}, errExit
}
