package service

import "route256/loms/internal/model"

// var _ loms.LomsServer = (*Service)(nil)

type Service struct {
	OrderRepository OrderRepo
	StockRepository StockRepo
}

type OrderRepo interface {
	Add(*model.Order) (model.OrderId, error)
	SetStatus(model.OrderId, model.OrderStatus) error
	GetById(model.OrderId) (*model.Order, error)
}

type StockRepo interface {
	Reserve(*model.OrderItem) error
	ReserveRemove(*model.OrderItem) error
	StockRemoveItem(*model.OrderItem) error
	GetStockItemBySku(uint32) (*model.StockItem, error)
}

func New(OrderRepository OrderRepo, StockRepository StockRepo) *Service {
	return &Service{
		OrderRepository,
		StockRepository,
	}
}
