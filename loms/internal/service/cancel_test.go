package service

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/service/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestOrderCancel(t *testing.T) {
	t.Parallel()
	order := &model.Order{
		User:   123,
		Status: model.StatusAwaitingPayment,
		Items: []*model.OrderItem{
			{
				Sku:   1,
				Count: 2,
			},
		},
	}

	tests := []struct {
		Name             string
		OrderId          model.OrderId
		Order            *model.Order
		WantGetByIdError error
		WantError        error
	}{
		{
			Name:             "Тест без  ошибок",
			OrderId:          model.OrderId(212),
			Order:            order,
			WantGetByIdError: nil,
			WantError:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			OrderMock := mocks.NewOrderRepoMock(ctrl)
			StockMock := mocks.NewStockRepoMock(ctrl)

			OrderMock.GetByIdMock.Expect(context.Background(), tt.OrderId).Return(tt.Order, tt.WantGetByIdError)
			OrderMock.SetStatusMock.Expect(context.Background(), tt.OrderId, model.StatusCancelled).Return(nil)
			StockMock.ReserveRemoveMock.Optional().Return(nil)

			NewService := New(OrderMock, StockMock)
			err := NewService.OrderCancel(context.Background(), tt.OrderId)
			assert.Equal(t, err, tt.WantError)
		})
	}
}
