package stock

import (
	"route256/loms/internal/model"
)

func (s *StockRepository) StockRemoveItem(order *model.OrderItem) error {
	item, ok := s.Repo[order.Sku]
	if !ok {
		return model.ErrSkuNoSuch
	}

	item.TotalCount -= uint64(order.Count)

	s.Repo[order.Sku] = item

	return nil
}
