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

	err = s.OrderRepository.OrderPay(ctx, orderId, order)
	if err != nil {
		return fmt.Errorf("s.OrderRepository.OrderPay %w", err)
	}

	s.Producer.MessagePush(&model.ProducerMessage{
		Topic:     string(model.TopicLomsOrderEvents),
		Partition: s.Producer.GetPartition(int32(orderId)),
		Value:     string(fmt.Sprintf("OrderId: %v Status: %v", orderId, model.StatusPayed)),
	})

	return nil
}
