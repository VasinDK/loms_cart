package delete_item

import (
	"context"
	"route256/cart/internal/service/item/delete_item/mock"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProductCart(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name      string
		cartId    int64
		sku       int64
		WantError error
	}{
		{
			Name:      "Обычные данные",
			cartId:    12,
			sku:       13,
			WantError: nil,
		},
		{
			Name:      "Нулевое значение cartId",
			cartId:    0,
			sku:       13,
			WantError: nil,
		},
		{
			Name:      "Нулевое значение sku",
			cartId:    12,
			sku:       0,
			WantError: nil,
		},
		{
			Name:      "Нулевое значение sku и cartId",
			cartId:    0,
			sku:       0,
			WantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			repositoryMock := mock.NewRepositoryMock(ctrl)
			repositoryMock.DeleteProductCartMock.Expect(context.Background(), tt.cartId, tt.sku).Return(tt.WantError)

			NewHandler := New(repositoryMock)
			err := NewHandler.DeleteProductCart(context.Background(), tt.cartId, tt.sku)
			assert.Equal(t, tt.WantError, err)
		})
	}
}
