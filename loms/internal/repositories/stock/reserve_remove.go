package stock

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"time"

	"github.com/jackc/pgx/v5"
)

// ReserveRemove - удаляет резерв
func (s *StockRepository) ReserveRemove(ctx context.Context, order *model.OrderItem) error {
	currentStockItem, err := s.GetItemsBySku(ctx, &[]uint32{order.Sku})
	if err != nil {
		return err
	}

	if len(*currentStockItem) <= 0 {
		return model.ErrSkuNoSuch
	}

	newReserved := int64((*currentStockItem)[0].Reserved - uint64(order.Count))
	if newReserved < 0 {
		return fmt.Errorf("in reserved %w", model.ErrSkuNotEnough)
	}

	const query = `
		UPDATE stocks
		SET reserved = @reserved
		WHERE sku = @sku
	`
	args := pgx.NamedArgs{
		"reserved": newReserved,
		"sku":      order.Sku,
	}

	start := time.Now()

	_, err = s.Conn.Exec(ctx, query, args)

	RequestDBTotal.WithLabelValues("UPDATE").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(err), "UPDATE").Observe(float64(time.Since(start).Seconds()))

	if err != nil {
		return err
	}

	return nil
}
