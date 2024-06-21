package stock

import (
	"context"
	"route256/loms/internal/model"

	"github.com/jackc/pgx/v5"
)

// GetItemBySku - получает стоки по sku
func (s *StockRepository) GetItemsBySku(ctx context.Context, sku *[]uint32) (*[]model.StockItem, error) {
	const query = `
		SELECT sku as Sku, total_count as TotalCount, reserved as Reserved
		FROM stocks
		WHERE sku = ANY($1)
	`
	rows, err := s.Conn.Query(ctx, query, sku)
	if err != nil {
		return nil, err
	}

	itemModel, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.StockItem])
	if err != nil {
		return nil, err
	}

	return &itemModel, nil
}
