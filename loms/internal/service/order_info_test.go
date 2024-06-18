package service

import (
	"context"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/internal/service/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestOrderInfo(t *testing.T) {
	t.Parallel()
	orderItem1 := &model.OrderItem{Sku: 1, Count: 2}

	order1 := &model.Order{
		User: 123,
		Items: []*model.OrderItem{
			orderItem1,
		},
	}

	tests := []struct {
		Name      string
		Order     *model.Order
		OrderId   model.OrderId
		WantError error
	}{
		{
			Name:      "Тест без ошибок",
			Order:     order1,
			OrderId:   model.OrderId(312),
			WantError: nil,
		},
		{
			Name:      "Тест c ошибкой",
			Order:     nil,
			OrderId:   model.OrderId(0),
			WantError: fmt.Errorf("%v, s.OrderRepository.GetById %w", "OrderInfo", model.ErrOrderNoSuch),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			OrderMock := mocks.NewOrderRepoMock(ctrl)
			StockMock := mocks.NewStockRepoMock(ctrl)

			OrderMock.GetByIdMock.Set(func(ctx context.Context, op1 model.OrderId) (*model.Order, error) {
				switch op1 {
				case model.OrderId(312):
					return order1, nil
				case model.OrderId(0):
					return nil, model.ErrOrderNoSuch
				default:
					return nil, nil
				}
			})

			NewService := New(OrderMock, StockMock)
			order, err := NewService.OrderInfo(context.Background(), tt.OrderId)
			assert.Equal(t, err, tt.WantError)
			assert.Equal(t, order, tt.Order)
		})
	}
}
