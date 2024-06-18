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

// OrderPay - покупка
func (h *Handlers) OrderPay(ctx context.Context, orderId *loms.OrderPayRequest) (*loms.OrderPayResponse, error) {
	const op = "OrderPay"

	err := h.service.OrderPay(ctx, model.OrderId(orderId.GetOrderId()))

	if errors.Is(err, model.ErrStatusNoAwaitingPayment) {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return &loms.OrderPayResponse{}, model.ErrStatusNoAwaitingPayment
	}

	if errors.Is(err, model.ErrSkuNoSuch) {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return &loms.OrderPayResponse{}, model.ErrSkuNoSuch
	}

	if err != nil {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return &loms.OrderPayResponse{}, status.Error(codes.Internal, "wtf")
	}

	return &loms.OrderPayResponse{}, nil
}
