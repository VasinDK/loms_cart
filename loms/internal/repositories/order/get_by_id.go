package order

import "route256/loms/internal/model"

func (o *OrderRepository) GetById(orderId model.OrderId) (*model.Order, error) {
	order, ok := o.Repo[orderId]
	if !ok {
		return nil, model.ErrOrderNoSuch
	}

	return order, nil
}
