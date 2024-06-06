package delete_item

import (
	"route256/cart/internal/service/item/delete_item/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gojuno/minimock/v3"
)

func TestDeleteProductCart(t *testing.T) {
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
	ctrl := minimock.NewController(t)
	repositoryMock := mock.NewRepositoryMock(ctrl)
	NewHandler := New(repositoryMock)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			repositoryMock.DeleteProductCartMock.Expect(tt.cartId, tt.sku).Return(tt.WantError)
			err := NewHandler.DeleteProductCart(tt.cartId, tt.sku)
			assert.Equal(t, tt.WantError, err)
		})
	}
}
