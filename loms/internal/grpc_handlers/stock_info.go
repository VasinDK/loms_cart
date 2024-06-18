package grpc_handlers

import (
	"context"
	"errors"
	"log/slog"
	"route256/loms/internal/model"
	"route256/loms/pkg/api/loms/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StocksInfo - информация о стоке
func (h *Handlers) StocksInfo(ctx context.Context, sku *loms.StocksInfoRequest) (*loms.StocksInfoResponse, error) {
	const op = "StocksInfo"

	count, err := h.service.StocksInfo(ctx, sku.GetSku())
	if errors.Is(err, model.ErrSkuNoSuch) {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return nil, model.ErrSkuNoSuch
	}

	if err != nil {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return nil, status.Error(codes.Internal, "Stock babah")
	}

	var cnt loms.StocksInfoResponse
	cnt.Count = count

	return &cnt, nil
}
