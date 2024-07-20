package service

import (
	"context"
	"route256/loms/internal/model"
)

// var _ loms.LomsServer = (*Service)(nil)

type Service struct {
	OrderRepository OrderRepo
	StockRepository StockRepo
	Producer        Producer
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

type Producer interface {
	MessagePush(*model.ProducerMessage)
	GetPartition(int32) int32
}

func New(OrderRepository OrderRepo, StockRepository StockRepo, Producer Producer) *Service {
	return &Service{
		OrderRepository,
		StockRepository,
		Producer,
	}
}
