package stock

import (
	"route256/loms/internal/model"
)

// Reserve - резервирует sku
func (s *StockRepository) Reserve(item *model.OrderItem) error {
	if _, ok := s.Repo[item.Sku]; !ok {
		return model.ErrSkuNoSuch
	}

	free := s.Repo[item.Sku].TotalCount - s.Repo[item.Sku].Reserved
	if (int64(free) - int64(item.Count)) < 0 {

		return model.ErrSkuNotEnough
	}

	stockItem := s.Repo[item.Sku]
	stockItem.Reserved += uint64(item.Count)
	s.Repo[item.Sku] = stockItem

	return nil
}
