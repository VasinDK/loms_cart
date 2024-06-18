package grpc_handlers

import (
	"context"
	"log/slog"
	"route256/loms/internal/model"
	"route256/loms/pkg/api/loms/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OrderInfo - инфа об ордере
func (h *Handlers) OrderInfo(ctx context.Context, OrderId *loms.OrderInfoRequest) (*loms.OrderInfoResponse, error) {
	const op = "OrderInfo"
	order, err := h.service.OrderInfo(ctx, model.OrderId(OrderId.GetOrderId()))
	if err != nil {
		slog.Error(op, "h.service.OrderInfo", err.Error())
		return nil, status.Error(codes.NotFound, "sorry nigga")
	}

	orderProto, err := h.RepackOrderToProto(order)
	if err != nil {
		slog.Error(op, "h.RepackOrderToProto", err.Error())
		return nil, status.Error(codes.NotFound, "sorry")
	}

	return orderProto, nil
}
