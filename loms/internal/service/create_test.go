package service

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/service/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	orderItem1 := &model.OrderItem{Sku: 1, Count: 2}
	orderItem2 := &model.OrderItem{Sku: 20, Count: 100}

	order1 := &model.Order{
		User: 123,
		Items: []*model.OrderItem{
			orderItem1,
		},
	}

	order2 := &model.Order{
		User: 423,
		Items: []*model.OrderItem{
			orderItem2,
		},
	}

	tests := []struct {
		Name         string
		Order        *model.Order
		OrderId      model.OrderId
		OrderItem    *model.OrderItem
		WantAddError error
		WantError    error
	}{
		{
			Name:         "Тест без ошибок",
			Order:        order1,
			OrderId:      model.OrderId(312),
			WantAddError: nil,
			WantError:    nil,
		},
		{
			Name:         "Тест c ошибкой",
			Order:        order2,
			OrderId:      model.OrderId(0),
			WantAddError: nil,
			WantError:    model.ErrSkuNoSuch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			OrderMock := mocks.NewOrderRepoMock(ctrl)
			StockMock := mocks.NewStockRepoMock(ctrl)

			OrderMock.AddMock.Expect(context.Background(), tt.Order).Return(tt.OrderId, tt.WantAddError)
			OrderMock.SetStatusMock.Optional().Return(nil)
			StockMock.ReserveMock.Optional().Set(func(ctx context.Context, op1 *model.OrderItem) error {
				switch op1 {
				case orderItem1:
					return nil
				case orderItem2:
					return model.ErrSkuNoSuch
				default:
					return nil
				}
			})

			NewService := New(OrderMock, StockMock)
			orderId, err := NewService.Create(context.Background(), tt.Order)
			assert.Equal(t, err, tt.WantError)
			assert.Equal(t, orderId, tt.OrderId)
		})
	}
}
