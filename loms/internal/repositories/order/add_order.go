package order

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"time"
)

// Add - добавляет ордер
func (o *OrderRepository) AddOrder(ctx context.Context, order *model.Order) (model.OrderId, error) {
	const query = `
		INSERT INTO orders (user_id, status) VALUES ($1, $2) RETURNING id;
	`
	var id int64

	start := time.Now()

	err := o.Conn.QueryRow(ctx, query, order.User, string(order.Status)).Scan(&id)

	RequestDBTotal.WithLabelValues("INSERT").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(err), "INSERT").Observe(float64(time.Since(start).Seconds()))

	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, model.ErrAddOrder
	}

	return model.OrderId(id), nil
}
