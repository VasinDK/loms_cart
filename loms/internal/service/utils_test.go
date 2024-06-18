package service

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/service/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestSetStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name      string
		OrderId   model.OrderId
		Status    model.OrderStatus
		Error     error
		WantError error
	}{
		{
			Name:      "Тест без ошибок",
			OrderId:   model.OrderId(312),
			Status:    model.StatusAwaitingPayment,
			Error:     nil,
			WantError: nil,
		},
		{
			Name:      "Тест c ошибкой",
			OrderId:   model.OrderId(0),
			Status:    model.StatusFailed,
			Error:     model.ErrOrderNoSuch,
			WantError: model.ErrOrderNoSuch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			OrderMock := mocks.NewOrderRepoMock(ctrl)
			StockMock := mocks.NewStockRepoMock(ctrl)

			OrderMock.SetStatusMock.Expect(context.Background(), tt.OrderId, tt.Status).Return(tt.Error)

			NewService := New(OrderMock, StockMock)
			err := NewService.SetStatus(context.Background(), tt.OrderId, tt.Status)
			assert.Equal(t, err, tt.WantError)
		})
	}
}
