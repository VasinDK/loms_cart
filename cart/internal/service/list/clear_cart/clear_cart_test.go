package clear_cart

import (
	"context"
	"route256/cart/internal/service/list/clear_cart/mock"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestClearCart(t *testing.T) {
	t.Parallel()
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

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			repositoryMock := mock.NewRepositoryMock(ctrl)
			repositoryMock.ClearCartMock.Expect(context.Background(), tt.cartId).Return(tt.WantError)

			NewHandler := New(repositoryMock)
			err := NewHandler.ClearCart(context.Background(), tt.cartId)
			assert.Equal(t, tt.WantError, err)
		})
	}
}
