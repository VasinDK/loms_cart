package service

import (
	"context"
	"route256/loms/internal/model"
)

// SetStatus - установка статуса
func (s *Service) SetStatus(ctx context.Context, orderId model.OrderId, status model.OrderStatus) error {
	err := s.OrderRepository.SetStatus(ctx, orderId, status)
	if err != nil {
		return err
	}

	return nil
}
