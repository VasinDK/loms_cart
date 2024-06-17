package service

import (
	"fmt"
	"route256/loms/internal/model"
)

// OrderCancel - удаляет ордер
func (s *Service) OrderCancel(orderId model.OrderId) error {
	order, err := s.OrderRepository.GetById(orderId)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.GetById %w", err)
	}

	err = s.OrderRepository.SetStatus(orderId, model.StatusCancelled)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.SetStatus %w", err)
	}

	for i := range order.Items {
		err = s.StockRepository.ReserveRemove(order.Items[i])
		if err != nil {
			return fmt.Errorf("s.StockRepository.ReserveRemove %w", err)
		}
	}

	return nil
}
