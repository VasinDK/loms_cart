package stock

import "route256/loms/internal/model"

func (s *StockRepository) GetStockItemBySku(sku uint32) (*model.StockItem, error) {
	item, ok := s.Repo[sku]
	if !ok {
		return nil, model.ErrSkuNoSuch
	}

	itemModel := model.StockItem(item)

	return &itemModel, nil
}
