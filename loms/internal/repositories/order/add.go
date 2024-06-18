package order

import (
	"context"
	"route256/loms/internal/model"
)

// Add - добавляет ордер
func (o *OrderRepository) Add(ctx context.Context, order *model.Order) (model.OrderId, error) {
	key := model.OrderId(len(o.Repo) + 1)
	o.Repo[key] = order

	return key, nil
}
