package service

import (
	"context"
	"fmt"
)

// StocksInfo - инфа о стоке
func (s *Service) StocksInfo(ctx context.Context, sku uint32) (uint64, error) {
	item, err := s.StockRepository.GetStockItemBySku(ctx, sku)
	if err != nil {
		return 0, fmt.Errorf("s.StockRepository.GetStockItemBySku %w", err)
	}

	remains := item.TotalCount - item.Reserved

	return remains, nil
}
