package stock

import (
	"context"
	"fmt"
	"route256/loms/internal/model"

	"github.com/jackc/pgx/v5"
)

// StockRemoveItem - удаляет элем.из стока
func (s *StockRepository) StockRemoveItem(ctx context.Context, order *model.OrderItem) error {
	currentStockItem, err := s.GetItemsBySku(ctx, &[]uint32{order.Sku})
	if err != nil {
		return err
	}

	if len(*currentStockItem) <= 0 {
		return model.ErrSkuNoSuch
	}

	newCount := int64((*currentStockItem)[0].TotalCount - uint64(order.Count))
	if newCount < 0 {
		return fmt.Errorf("in stock %w", model.ErrSkuNotEnough)
	}

	const query = `
		UPDATE stocks
		SET total_count = @total_count
		WHERE sku = @sku
	`
	args := pgx.NamedArgs{
		"total_count": newCount,
		"sku":         order.Sku,
	}

	_, err = s.Conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}
