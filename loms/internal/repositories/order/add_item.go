package order

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

// Add - добавляет ордер
func (o *OrderRepository) AddItem(ctx context.Context, order *model.Order, orderId model.OrderId) error {
	const query = `
		INSERT INTO items_order (sku, count, order_id) VALUES (@sku, @count, @orderId);
	`
	batch := &pgx.Batch{}

	for i := range order.Items {
		args := pgx.NamedArgs{
			"sku":     order.Items[i].Sku,
			"count":   order.Items[i].Count,
			"orderId": orderId,
		}

		batch.Queue(query, args)
	}

	start := time.Now()

	res := o.Conn.SendBatch(ctx, batch)
	defer res.Close()

	var errForLabel error
	var once sync.Once

	for _, v := range order.Items {
		_, err := res.Exec()
		if err != nil {
			once.Do(func() { errForLabel = err })
			return fmt.Errorf("in %v error: %w", v, err)
		}
	}

	RequestDBTotal.WithLabelValues("INSERT").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(errForLabel), "INSERT").Observe(float64(time.Since(start).Seconds()))

	return nil
}
