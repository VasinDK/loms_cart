package stock

import (
	"context"
	"route256/loms/internal/model"
)

// GetStockItemBySku - получает стоки по sku
func (s *StockRepository) GetStockItemBySku(ctx context.Context, sku uint32) (*model.StockItem, error) {
	item, ok := s.Repo[sku]
	if !ok {
		return nil, model.ErrSkuNoSuch
	}

	itemModel := model.StockItem(item)

	return &itemModel, nil
}
