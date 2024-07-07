package order

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"time"
)

// GetById - получает ордер по id
func (o *OrderRepository) GetById(ctx context.Context, orderId model.OrderId) (*model.Order, error) {
	const query = `
		SELECT orders.user_id, orders.status, items_order.sku, items_order.count
		FROM orders
		JOIN items_order ON orders.id = items_order.order_id
		WHERE orders.id=$1
	`
	start := time.Now()

	rows, err := o.Conn.Query(ctx, query, orderId)

	RequestDBTotal.WithLabelValues("SELECT").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(err), "SELECT").Observe(float64(time.Since(start).Seconds()))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	order := model.Order{}
	var user int64
	var status string

	for rows.Next() {
		item := model.OrderItem{}
		rows.Scan(&user, &status, &item.Sku, &item.Count)
		order.Items = append(order.Items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	order.User = user
	order.Status = model.OrderStatus(status)

	return &order, nil
}
