package order

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"time"
)

// Add - добавляет ордер
func (o *OrderRepository) AddOrder(ctx context.Context, order *model.Order) (model.OrderId, error) {
	shIndex := o.Sm.GetShardIndexFromID(order.User)
	Conn, err := o.Sm.Pick(shIndex)
	if err != nil {
		return 0, fmt.Errorf("o.Sm.Pick %w", err)
	}

	const query = `
		INSERT INTO orders (id, user_id, status) 
		VALUES (nextval('order_id_manual_seq') + $1, $2) RETURNING id;
	`
	var id model.OrderId

	start := time.Now()

	err = Conn.QueryRow(ctx, query, shIndex, order.User, string(order.Status)).Scan(&id)

	RequestDBTotal.WithLabelValues("INSERT").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(err), "INSERT").Observe(float64(time.Since(start).Seconds()))

	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, model.ErrAddOrder
	}

	return id, nil
}
