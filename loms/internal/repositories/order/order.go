package order

import "route256/loms/internal/model"

type OrderRepository struct {
	Repo map[model.OrderId]*model.Order
}

func New() *OrderRepository {
	return &OrderRepository{
		Repo: make(map[model.OrderId]*model.Order),
	}
}
