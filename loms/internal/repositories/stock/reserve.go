package stock

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/pkg/statuses"
	"sync"
	"time"

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

	start := time.Now()

	res := s.Conn.SendBatch(ctx, batch)
	defer res.Close()

	var errForLabel error
	var once sync.Once

	for _, v := range *stockItems {
		_, err := res.Exec()
		if err != nil {
			once.Do(func() { errForLabel = err })
			return fmt.Errorf("in %v error: %w", v, err)
		}
	}

	RequestDBTotal.WithLabelValues("UPDATE").Inc()
	RequestTimeStatusCategoryBD.WithLabelValues(statuses.GetCodePG(errForLabel), "UPDATE").Observe(float64(time.Since(start).Seconds()))

	return nil
}
