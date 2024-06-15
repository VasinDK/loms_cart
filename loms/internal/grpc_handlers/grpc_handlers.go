package grpc_handlers

import (
	"route256/loms/internal/model"
	"route256/loms/pkg/api/loms/v1"
)

type Handlers struct {
	loms.UnimplementedLomsServer
	service Service
}

type Service interface {
	Create(*model.Order) (model.OrderId, error)
	OrderInfo(model.OrderId) (*model.Order, error)
	OrderPay(model.OrderId) error
	OrderCancel(model.OrderId) error
	StocksInfo(uint32) (uint64, error)
}

func New(service Service) *Handlers {
	return &Handlers{
		service: service,
	}
}
