package service

import (
	"fmt"
	"route256/loms/internal/model"
)

func (s *Service) OrderPay(orderId model.OrderId) error {
	order, err := s.OrderRepository.GetById(orderId)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.GetById %w", err)
	}

	if order.Status != model.StatusAwaitingPayment {
		return model.ErrStatusNoAwaitingPayment
	}

	err = s.OrderRepository.SetStatus(orderId, model.StatusPayed)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.SetStatus %w", err)
	}

	for i := range order.Items {
		err = s.StockRepository.ReserveRemove(order.Items[i])
		if err != nil {
			return fmt.Errorf("s.StockRepository.ReserveRemove %w", err)
		}

		err = s.StockRepository.StockRemoveItem(order.Items[i])
		if err != nil {
			return fmt.Errorf("s.StockRepository.StockRemoveItem %w", err)
		}

	}

	return nil
}
