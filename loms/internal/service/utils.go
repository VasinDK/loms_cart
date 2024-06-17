package service

import "route256/loms/internal/model"

// SetStatus - установка статуса
func (s *Service) SetStatus(orderId model.OrderId, status model.OrderStatus) error {
	err := s.OrderRepository.SetStatus(orderId, status)
	if err != nil {
		return err
	}

	return nil
}
