package order

import (
	"context"
	"route256/loms/internal/model"
)

// SetStatus - устанавливает статус ордера
func (o *OrderRepository) SetStatus(ctx context.Context, orderId model.OrderId, status model.OrderStatus) error {
	if _, ok := o.Repo[orderId]; !ok {
		return model.ErrOrderNoSuch
	}

	o.Repo[orderId].Status = status

	return nil
}
