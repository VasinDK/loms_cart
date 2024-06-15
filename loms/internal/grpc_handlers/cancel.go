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

func (h *Handlers) OrderCancel(ctx context.Context, orderId *loms.OrderId) (*emptypb.Empty, error) {
	const op = "OrderCancel"

	err := h.service.OrderCancel(model.OrderId(orderId.GetOrderId()))
	if errors.Is(err, model.ErrSkuNoSuch) {
		slog.Error(op, "h.service.OrderCancel", err.Error())
		return &emptypb.Empty{}, model.ErrSkuNoSuch
	}

	if err != nil {
		slog.Error(op, "h.service.OrderCancel", err.Error())
		return &emptypb.Empty{}, status.Error(codes.Internal, "wtf")
	}

	return &emptypb.Empty{}, nil
}
