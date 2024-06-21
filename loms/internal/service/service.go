package service

import (
	"context"
	"route256/loms/internal/model"
)

// var _ loms.LomsServer = (*Service)(nil)

type Service struct {
	OrderRepository OrderRepo
	StockRepository StockRepo
}

type OrderRepo interface {
	AddOrder(context.Context, *model.Order) (model.OrderId, error)
	AddItem(context.Context, *model.Order, model.OrderId) error
	SetStatus(context.Context, model.OrderId, model.OrderStatus) error
	GetById(context.Context, model.OrderId) (*model.Order, error)
	OrderPay(context.Context, model.OrderId, *model.Order) error
}

type StockRepo interface {
	Reserve(context.Context, *[]model.StockItem) error
	ReserveRemove(context.Context, *model.OrderItem) error
	StockRemoveItem(context.Context, *model.OrderItem) error
	GetItemsBySku(context.Context, *[]uint32) (*[]model.StockItem, error)
}

func New(OrderRepository OrderRepo, StockRepository StockRepo) *Service {
	return &Service{
		OrderRepository,
		StockRepository,
	}
}
