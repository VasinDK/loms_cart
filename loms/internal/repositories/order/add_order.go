package order

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"strconv"
	"time"
)

// Add - добавляет ордер
func (o *OrderRepository) AddOrder(ctx context.Context, order *model.Order) (model.OrderId, error) {
	shIndex, err := o.Sm.GetShardIndex(strconv.FormatInt(order.User, 10))
	if err != nil {
		return 0, fmt.Errorf("o.Sm.GetShardIndex %w", err)
	}

	Conn, err := o.Sm.Pick(int(shIndex))
	if err != nil {
		return 0, fmt.Errorf("o.Sm.Pick %w", err)
	}

	const query = `
		INSERT INTO orders (id, user_id, status) 
		VALUES (nextval('order_id_manual_seq') + $1, $2, $3) RETURNING id;
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
