package grpc_handlers

import (
	"context"
	"log/slog"
	"route256/loms/pkg/api/loms/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handlers) OrderCreate(ctx context.Context, order *loms.OrderCreateRequest) (*loms.OrderId, error) {
	const op = "OrderCreate"

	orderModel, err := h.RepackOrderToModel(order)
	if err != nil {
		slog.Error(op, "h.RepackOrderInOrderModel", err.Error())
		return nil, status.Error(codes.FailedPrecondition, "so")
	}

	orderIdModel, err := h.service.Create(orderModel)
	if err != nil {
		slog.Error(op, "h.service.OrderCreate", err.Error())
		return nil, status.Error(codes.FailedPrecondition, "I am sorry")
	}

	return h.RepackOrderIdToProto(orderIdModel), nil
}
