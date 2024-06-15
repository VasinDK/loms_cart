package order

import (
	"route256/loms/internal/model"
)

func (o *OrderRepository) SetStatus(orderId model.OrderId, status model.OrderStatus) error {
	o.Repo[orderId].Status = status
	return nil
}
