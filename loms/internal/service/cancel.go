package service

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
)

// OrderCancel - удаляет ордер
func (s *Service) OrderCancel(ctx context.Context, orderId model.OrderId) error {
	order, err := s.OrderRepository.GetById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.GetById %w", err)
	}

	err = s.OrderRepository.SetStatus(ctx, orderId, model.StatusCancelled)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.SetStatus %w", err)
	}

	for i := range order.Items {
		err = s.StockRepository.ReserveRemove(ctx, order.Items[i])
		if err != nil {
			return fmt.Errorf("s.StockRepository.ReserveRemove %w", err)
		}
	}

	return nil
}
