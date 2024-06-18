package order

import (
	"context"
	"route256/loms/internal/model"
)

// GetById - получает ордер по id
func (o *OrderRepository) GetById(ctx context.Context, orderId model.OrderId) (*model.Order, error) {
	order, ok := o.Repo[orderId]
	if !ok {
		return nil, model.ErrOrderNoSuch
	}

	return order, nil
}
