package service

import (
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/internal/service/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestStocksInfo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name            string
		Sku             uint32
		Item            *model.StockItem
		WantGetStockErr error
		WantError       error
		Remains         uint64
	}{
		{
			Name: "Тест без ошибок",
			Sku:  121,
			Item: &model.StockItem{
				TotalCount: 100,
				Reserved:   60,
			},
			WantGetStockErr: nil,
			WantError:       nil,
			Remains:         40,
		},
		{
			Name: "Тест c ошибкой",
			Sku:  12,
			Item: &model.StockItem{
				TotalCount: 100,
				Reserved:   30,
			},
			WantGetStockErr: model.ErrSkuNoSuch,
			WantError:       fmt.Errorf("s.StockRepository.GetStockItemBySku %w", model.ErrSkuNoSuch),
			Remains:         0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			OrderMock := mocks.NewOrderRepoMock(ctrl)
			StockMock := mocks.NewStockRepoMock(ctrl)

			StockMock.GetStockItemBySkuMock.Expect(tt.Sku).Return(tt.Item, tt.WantGetStockErr)

			NewService := New(OrderMock, StockMock)
			remains, err := NewService.StocksInfo(tt.Sku)
			assert.Equal(t, err, tt.WantError)
			assert.Equal(t, remains, tt.Remains)
		})
	}
}
