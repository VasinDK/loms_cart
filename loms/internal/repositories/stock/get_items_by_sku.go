package stock

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"time"

	"github.com/jackc/pgx/v5"
)

// GetItemBySku - получает стоки по sku
func (s *StockRepository) GetItemsBySku(ctx context.Context, sku *[]uint32) (*[]model.StockItem, error) {
	Conn, err := s.Sm.Pick(s.Sm.GetMainShard())
	if err != nil {
		return nil, fmt.Errorf("s.Sm.Pick %w", err)
	}

	const query = `
		SELECT sku as Sku, total_count as TotalCount, reserved as Reserved
		FROM stocks
		WHERE sku = ANY($1)
	`
	start := time.Now()

	rows, err := Conn.Query(ctx, query, sku)

	RequestDBTotal.WithLabelValues("SELECT").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(err), "SELECT").Observe(time.Since(start).Seconds())

	if err != nil {
		return nil, err
	}

	itemModel, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.StockItem])
	if err != nil {
		return nil, err
	}

	return &itemModel, nil
}
