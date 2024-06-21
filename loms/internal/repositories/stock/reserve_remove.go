package stock

import (
	"context"
	"fmt"
	"route256/loms/internal/model"

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

	_, err = s.Conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}
