package service

import "fmt"

// StocksInfo - инфа о стоке
func (s *Service) StocksInfo(sku uint32) (uint64, error) {
	item, err := s.StockRepository.GetStockItemBySku(sku)
	if err != nil {
		return 0, fmt.Errorf("s.StockRepository.GetStockItemBySku %w", err)
	}

	remains := item.TotalCount - item.Reserved

	return remains, nil
}
