package service

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
)

// StocksInfo - инфа о стоке
func (s *Service) StocksInfo(ctx context.Context, sku uint32) (uint64, error) {
	items, err := s.StockRepository.GetItemsBySku(ctx, &[]uint32{sku})
	if err != nil {
		return 0, fmt.Errorf("s.StockRepository.GetItemsBySku %w", err)
	}

	if len(*items) > 1 {
		return 0, fmt.Errorf("s.StockRepository.GetItemsBySku %w", model.ErrDuplicateSku)
	}

	if len(*items) == 0 {
		return 0, fmt.Errorf("s.StockRepository.GetItemsBySku %w", model.ErrSkuNoSuch)
	}

	remains := (*items)[0].TotalCount - (*items)[0].Reserved

	return remains, nil
}
