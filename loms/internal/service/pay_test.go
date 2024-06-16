package service

import (
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/internal/service/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestOrderPay(t *testing.T) {
	t.Parallel()
	orderItem1 := &model.OrderItem{Sku: 1, Count: 2}
	orderItem2 := &model.OrderItem{Sku: 20, Count: 100}

	order1 := &model.Order{
		User:   123,
		Status: model.StatusAwaitingPayment,
		Items: []*model.OrderItem{
			orderItem1,
		},
	}

	order2 := &model.Order{
		User:   423,
		Status: model.StatusFailed,
		Items: []*model.OrderItem{
			orderItem2,
		},
	}

	tests := []struct {
		Name             string
		Order            *model.Order
		OrderId          model.OrderId
		OrderItem        *model.OrderItem
		WantGetByIdError error
		WantError        error
	}{
		{
			Name:             "Без ошибок",
			Order:            order1,
			OrderId:          model.OrderId(312),
			WantGetByIdError: nil,
			WantError:        nil,
		},
		{
			Name:             "Ошибка. Нет такого OrderId",
			Order:            nil,
			OrderId:          model.OrderId(312),
			WantGetByIdError: model.ErrOrderNoSuch,
			WantError:        fmt.Errorf("s.OrderRepository.GetById %w", model.ErrOrderNoSuch),
		},
		{
			Name:             "Ошибка. Статус заявки не соответствует",
			Order:            order2,
			OrderId:          model.OrderId(423),
			WantGetByIdError: nil,
			WantError:        model.ErrStatusNoAwaitingPayment,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			OrderMock := mocks.NewOrderRepoMock(ctrl)
			StockMock := mocks.NewStockRepoMock(ctrl)

			OrderMock.GetByIdMock.Expect(tt.OrderId).Return(tt.Order, tt.WantGetByIdError)
			OrderMock.SetStatusMock.Optional().Return(nil)
			StockMock.ReserveRemoveMock.Optional().Return(nil)
			StockMock.StockRemoveItemMock.Optional().Return(nil)

			NewService := New(OrderMock, StockMock)
			err := NewService.OrderPay(tt.OrderId)
			assert.Equal(t, err, tt.WantError)
		})
	}
}
