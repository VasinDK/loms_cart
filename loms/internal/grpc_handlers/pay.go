package grpc_handlers

import (
	"context"
	"errors"
	"log/slog"
	"route256/loms/internal/model"
	"route256/loms/pkg/api/loms/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// OrderPay - покупка
func (h *Handlers) OrderPay(ctx context.Context, orderId *loms.OrderId) (*emptypb.Empty, error) {
	const op = "OrderPay"

	err := h.service.OrderPay(model.OrderId(orderId.GetOrderId()))

	if errors.Is(err, model.ErrStatusNoAwaitingPayment) {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return &emptypb.Empty{}, model.ErrStatusNoAwaitingPayment
	}

	if errors.Is(err, model.ErrSkuNoSuch) {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return &emptypb.Empty{}, model.ErrSkuNoSuch
	}

	if err != nil {
		slog.Error(op, "h.service.OrderPay", err.Error())
		return &emptypb.Empty{}, status.Error(codes.Internal, "wtf")
	}

	return &emptypb.Empty{}, nil
}
