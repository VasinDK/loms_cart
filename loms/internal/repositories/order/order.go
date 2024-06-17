package order

import "route256/loms/internal/model"

type OrderRepository struct {
	Repo map[model.OrderId]*model.Order
}

// New - создает OrderRepository
func New() *OrderRepository {
	return &OrderRepository{
		Repo: make(map[model.OrderId]*model.Order),
	}
}
