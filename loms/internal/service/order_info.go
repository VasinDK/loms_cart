package service

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
)

// OrderInfo - получает ордер по id
func (s *Service) OrderInfo(ctx context.Context, orderId model.OrderId) (*model.Order, error) {
	const op = "OrderInfo"
	order, err := s.OrderRepository.GetById(ctx, orderId)
	if err != nil {
		return nil, fmt.Errorf("%v, s.OrderRepository.GetById %w", op, err)
	}

	if len(order.Items) <= 0 {
		return nil, model.ErrOrderNoSuch
	}

	return order, nil
}
