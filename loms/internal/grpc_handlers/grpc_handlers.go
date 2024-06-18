package grpc_handlers

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/pkg/api/loms/v1"
)

type Handlers struct {
	loms.UnimplementedLomsServer
	service Service
}

type Service interface {
	Create(context.Context, *model.Order) (model.OrderId, error)
	OrderInfo(context.Context, model.OrderId) (*model.Order, error)
	OrderPay(context.Context, model.OrderId) error
	OrderCancel(context.Context, model.OrderId) error
	StocksInfo(context.Context, uint32) (uint64, error)
}

func New(service Service) *Handlers {
	return &Handlers{
		service: service,
	}
}
