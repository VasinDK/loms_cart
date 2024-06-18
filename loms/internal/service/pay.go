package service

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
)

// OrderPay - оплата ордера по id. Проверка и перевод в статус
func (s *Service) OrderPay(ctx context.Context, orderId model.OrderId) error {
	order, err := s.OrderRepository.GetById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.GetById %w", err)
	}

	if order.Status != model.StatusAwaitingPayment {
		return model.ErrStatusNoAwaitingPayment
	}

	err = s.OrderRepository.SetStatus(ctx, orderId, model.StatusPayed)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.SetStatus %w", err)
	}

	for i := range order.Items {
		err = s.StockRepository.ReserveRemove(ctx, order.Items[i])
		if err != nil {
			return fmt.Errorf("s.StockRepository.ReserveRemove %w", err)
		}

		err = s.StockRepository.StockRemoveItem(ctx, order.Items[i])
		if err != nil {
			return fmt.Errorf("s.StockRepository.StockRemoveItem %w", err)
		}

	}

	return nil
}
