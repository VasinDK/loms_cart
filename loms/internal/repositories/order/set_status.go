package order

import (
	"context"
	"route256/loms/internal/model"
)

// SetStatus - устанавливает статус ордера
func (o *OrderRepository) SetStatus(ctx context.Context, orderId model.OrderId, status model.OrderStatus) error {
	const query = `
		UPDATE orders
		SET status = $2
		WHERE id = $1
	`
	_, err := o.Conn.Exec(ctx, query, orderId, status)
	if err != nil {
		return err
	}

	return nil
}
