package stock

import (
	"route256/loms/internal/model"
)

// ReserveRemove - удаляет резерв
func (s *StockRepository) ReserveRemove(order *model.OrderItem) error {
	item, ok := s.Repo[order.Sku]
	if !ok {
		return model.ErrSkuNoSuch
	}

	item.Reserved -= uint64(order.Count)

	s.Repo[order.Sku] = item

	return nil
}
