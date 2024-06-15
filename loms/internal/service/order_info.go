package service

import (
	"fmt"
	"route256/loms/internal/model"
)

func (s *Service) OrderInfo(orderId model.OrderId) (*model.Order, error) {
	const op = "OrderInfo"
	order, err := s.OrderRepository.GetById(orderId)

	if err != nil {
		return nil, fmt.Errorf("%v, s.OrderRepository.GetById %w", op, err)
	}

	return order, nil
}
