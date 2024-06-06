package clear_cart

import (
	"route256/cart/internal/service/list/clear_cart/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gojuno/minimock/v3"
)

func TestClearCart(t *testing.T) {
	tests := []struct {
		Name      string
		cartId    int64
		WantError error
	}{
		{
			Name:      "Типичное значение",
			cartId:    12,
			WantError: nil,
		},
		{
			Name:      "Ноль",
			cartId:    0,
			WantError: nil,
		},
	}
	ctrl := minimock.NewController(t)
	repositoryMock := mock.NewRepositoryMock(ctrl)
	NewHandler := New(repositoryMock)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			repositoryMock.ClearCartMock.Expect(tt.cartId).Return(tt.WantError)
			err := NewHandler.ClearCart(tt.cartId)
			assert.Equal(t, tt.WantError, err)
		})
	}
}
