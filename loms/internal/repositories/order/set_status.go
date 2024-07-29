package order

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"time"
)

// SetStatus - устанавливает статус ордера
func (o *OrderRepository) SetStatus(ctx context.Context, orderId model.OrderId, orderStatus model.OrderStatus) error {
	shIndex := o.Sm.GetShardIndexFromID(int64(orderId))
	Conn, err := o.Sm.Pick(shIndex)
	if err != nil {
		return fmt.Errorf("o.Sm.Pick %w", err)
	}

	const query = `
		UPDATE orders
		SET status = $2
		WHERE id = $1
	`
	start := time.Now()

	_, err = Conn.Exec(ctx, query, orderId, orderStatus)

	RequestDBTotal.WithLabelValues("UPDATE").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(err), "UPDATE").Observe(float64(time.Since(start).Seconds()))

	if err != nil {
		return err
	}

	return nil
}
