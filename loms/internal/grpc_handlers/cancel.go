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

// OrderCancel - отменяет ордер
func (h *Handlers) OrderCancel(ctx context.Context, orderId *loms.OrderCancelRequest) (*loms.OrderCancelResponse, error) {
	const op = "OrderCancel"
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

	err := h.service.OrderCancel(ctx, model.OrderId(orderId.GetOrderId()))
	if errors.Is(err, model.ErrSkuNoSuch) {
		logger.Errorw(ctx, op, "h.service.OrderCancel", "err", err)
		errExit = status.Error(codes.InvalidArgument, model.ErrSkuNoSuch.Error())
		return &loms.OrderCancelResponse{}, errExit
	}

	if err != nil {
		logger.Errorw(ctx, op, "h.service.OrderCancel", "err", err)
		errExit = status.Error(codes.Internal, "wtf")
		return &loms.OrderCancelResponse{}, errExit
	}

	return &loms.OrderCancelResponse{}, errExit
}
