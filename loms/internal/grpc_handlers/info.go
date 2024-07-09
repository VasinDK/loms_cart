package grpc_handlers

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/pkg/logger"
	"route256/loms/pkg/api/loms/v1"
	"route256/loms/pkg/statuses"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OrderInfo - инфа об ордере
func (h *Handlers) OrderInfo(ctx context.Context, OrderId *loms.OrderInfoRequest) (*loms.OrderInfoResponse, error) {
	const op = "OrderInfo"
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

	order, err := h.service.OrderInfo(ctx, model.OrderId(OrderId.GetOrderId()))
	if err != nil {
		logger.Errorw(ctx, op, "h.service.OrderInfo", "err", err)
		errExit = status.Error(codes.NotFound, "sorry nigga")
		return nil, errExit
	}

	orderProto, err := h.RepackOrderToProto(order)
	if err != nil {
		logger.Errorw(ctx, op, "h.RepackOrderToProto", "err", err)
		errExit = status.Error(codes.NotFound, "sorry")
		return nil, errExit
	}

	return orderProto, errExit
}
