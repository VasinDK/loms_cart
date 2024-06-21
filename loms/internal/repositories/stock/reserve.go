package stock

import (
	"context"
	"fmt"
	"route256/loms/internal/model"

	"github.com/jackc/pgx/v5"
)

// Reserve - резервирует sku
func (s *StockRepository) Reserve(ctx context.Context, stockItems *[]model.StockItem) error {
	query := `
		UPDATE stocks
		SET reserved = @reserved
		WHERE sku = @sku;
	`
	batch := &pgx.Batch{}
	
	for i := range *stockItems {
		args := pgx.NamedArgs{
			"reserved": (*stockItems)[i].Reserved,
			"sku":      (*stockItems)[i].Sku,
		}

		batch.Queue(query, args)
	}

	res := s.Conn.SendBatch(ctx, batch)
	defer res.Close()

	for _, v := range *stockItems {
		_, err := res.Exec()
		if err != nil {
			return fmt.Errorf("in %v error: %w", v, err)
		}
	}

	return nil
}
