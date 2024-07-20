package service

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/internal/pkg/logger"
)

// Create - создает ордер
func (s *Service) Create(ctx context.Context, order *model.Order) (model.OrderId, error) {
	const op = "Service.OrderCreate"

	order.Status = model.StatusNew
	orderId, err := s.OrderRepository.AddOrder(ctx, order)
	if err != nil {
		return 0, fmt.Errorf("%v s.OrderRepository.Add %w", op, err)
	}

	s.Producer.MessagePush(&model.ProducerMessage{
		Topic:     string(model.TopicLomsOrderEvents),
		Partition: s.Producer.GetPartition(int32(orderId)),
		Value:     string(fmt.Sprintf("OrderId: %v Status: %v", orderId, order.Status)),
	})

	err = s.OrderRepository.AddItem(ctx, order, orderId)
	if err != nil {
		return 0, fmt.Errorf("%v s.OrderRepository.Add %w", op, err)
	}

	skus := make([]uint32, len(order.Items))
	skusMap := make(map[uint32]uint16)

	for i, item := range order.Items {
		skus[i] = item.Sku
		skusMap[item.Sku] = item.Count
	}

	items, err := s.StockRepository.GetItemsBySku(ctx, &skus)
	if err != nil {
		return 0, fmt.Errorf("%v s.StockRepository.GetItemsBySku %w", op, err)
	}

	newStockItem := make([]model.StockItem, len(*items))

	for i := range *items {
		count, ok := skusMap[(*items)[i].Sku]
		if !ok {
			s.setStatus(ctx, op, orderId, model.StatusFailed)
			return 0, fmt.Errorf("%v s.StockRepository.GetItemsBySku %w", op, model.ErrDuplicateSku)
		}

		free := (*items)[i].TotalCount - (*items)[i].Reserved

		if (int64(free) - int64(count)) < 0 {
			s.setStatus(ctx, op, orderId, model.StatusFailed)
			return 0, fmt.Errorf("%v s.StockRepository.GetItemsBySku %w", op, model.ErrSkuNotEnough)
		}

		newStockItem[i] = model.StockItem{
			Sku:      (*items)[i].Sku,
			Reserved: (*items)[i].Reserved + uint64(count),
		}

		delete(skusMap, (*items)[i].Sku)
	}

	if len(skusMap) != 0 {
		s.setStatus(ctx, op, orderId, model.StatusFailed)
		return 0, fmt.Errorf("%v s.StockRepository.GetItemsBySku %w", op, model.ErrSkuNoSuch)
	}

	err = s.StockRepository.Reserve(ctx, &newStockItem)
	if err != nil {
		s.setStatus(ctx, op, orderId, model.StatusFailed)
		return 0, fmt.Errorf("%v s.StockRepository.Reserve %v %w", op, err.Error(), model.ErrSkuNotEnough)
	}

	// тут SetStatus основная функция поэтому делается проверка и возврат ошибки
	err = s.OrderRepository.SetStatus(ctx, orderId, model.StatusAwaitingPayment)
	if err != nil {
		return 0, model.ErrAddStatus
	}
	s.Producer.MessagePush(&model.ProducerMessage{
		Topic:     string(model.TopicLomsOrderEvents),
		Partition: s.Producer.GetPartition(int32(orderId)),
		Value:     string(fmt.Sprintf("OrderId: %v Status: %v", orderId, model.StatusAwaitingPayment)),
	})

	return orderId, nil
}

// setStatus - это дополнительная функция которая под капотом использует основную SetStatus.
// Эта функция запускается после возникновения основной ошибки. Возвращается основная ошибка,
// но при этом проверяется ошибка установки статуса. Если такая появляется она просто логируется
func (s *Service) setStatus(ctx context.Context, op string, orderId model.OrderId, Status model.OrderStatus) {
	errChangStatus := s.OrderRepository.SetStatus(ctx, orderId, Status)
	if errChangStatus != nil {
		logger.Errorw(ctx, op, "s.OrderRepository.SetStatus", "errChangStatus", errChangStatus)
	}

	s.Producer.MessagePush(&model.ProducerMessage{
		Topic:     string(model.TopicLomsOrderEvents),
		Partition: s.Producer.GetPartition(int32(orderId)),
		Value:     string(fmt.Sprintf("OrderId: %v Status: %v", orderId, Status)),
	})
}
