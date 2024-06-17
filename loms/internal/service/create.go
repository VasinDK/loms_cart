package service

import (
	"fmt"
	"route256/loms/internal/model"
)

// Create - создает ордер
func (s *Service) Create(order *model.Order) (model.OrderId, error) {
	const op = "Service.OrderCreate"

	order.Status = model.StatusNew
	orderId, err := s.OrderRepository.Add(order)
	if err != nil {
		return 0, fmt.Errorf("%v s.OrderRepository.Add %w", op, err)
	}

	for _, item := range order.Items {
		err = s.StockRepository.Reserve(item)

		if err != nil {
			errChangStatus := s.OrderRepository.SetStatus(orderId, model.StatusFailed)
			if errChangStatus != nil {
				return 0, fmt.Errorf("%v, %w", errChangStatus.Error(), err)
			}

			return 0, err
		}
	}

	err = s.OrderRepository.SetStatus(orderId, model.StatusAwaitingPayment)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}
