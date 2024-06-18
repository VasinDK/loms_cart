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

// OrderCancel - отменяет ордер
func (h *Handlers) OrderCancel(ctx context.Context, orderId *loms.OrderCancelRequest) (*loms.OrderCancelResponse, error) {
	const op = "OrderCancel"

	err := h.service.OrderCancel(ctx, model.OrderId(orderId.GetOrderId()))
	if errors.Is(err, model.ErrSkuNoSuch) {
		slog.Error(op, "h.service.OrderCancel", err.Error())
		return &loms.OrderCancelResponse{}, model.ErrSkuNoSuch
	}

	if err != nil {
		slog.Error(op, "h.service.OrderCancel", err.Error())
		return &loms.OrderCancelResponse{}, status.Error(codes.Internal, "wtf")
	}

	return &loms.OrderCancelResponse{}, nil
}
