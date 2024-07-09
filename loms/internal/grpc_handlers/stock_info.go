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

// StocksInfo - информация о стоке
func (h *Handlers) StocksInfo(ctx context.Context, sku *loms.StocksInfoRequest) (*loms.StocksInfoResponse, error) {
	const op = "StocksInfo"
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

	count, err := h.service.StocksInfo(ctx, sku.GetSku())
	if errors.Is(err, model.ErrSkuNoSuch) {
		logger.Errorw(ctx, op, "h.service.OrderPay", "err", err)
		errExit = status.Error(codes.InvalidArgument, model.ErrSkuNoSuch.Error())
		return nil, errExit
	}

	if err != nil {
		logger.Errorw(ctx, op, "h.service.OrderPay", "err", err)
		errExit = status.Error(codes.Internal, "Stock babah")
		return nil, errExit
	}

	var cnt loms.StocksInfoResponse
	cnt.Count = count

	return &cnt, errExit
}
