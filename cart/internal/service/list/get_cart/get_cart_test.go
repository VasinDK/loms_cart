package get_cart

import (
	"route256/cart/internal/model"
	"route256/cart/internal/service/list/get_cart/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gojuno/minimock/v3"
)

func TestGetCart(t *testing.T) {
	CartProd := map[int64]*model.Product{
		123: &model.Product{
			SKU:   123,
			Count: 2,
		},
		124: &model.Product{
			SKU:   124,
			Count: 1,
		},
	}
	CartProd0 := map[int64]*model.Product{}

	CheckSKUResp1 := &model.Product{ // SKU:   123
		Name:  "Чай",
		Price: 13,
	}

	CheckSKUResp2 := &model.Product{ // SKU:   124
		Name:  "Кофе",
		Price: 130,
	}
	// CheckSKUResp0 := &model.Product{}

	tests := []struct {
		Name        string
		CartId      int64
		CartProd    map[int64]*model.Product
		CheckSKUReq int64
		WantError error
		WantRes   uint32
	}{
		{
			Name:        "Типичное значение",
			CartId:      12,
			CartProd:    CartProd,
			CheckSKUReq: 123,
			WantError: nil,
			WantRes:   156,
		},
		{
			Name:        "Пустая корзина",
			CartId:      2,
			CartProd:    CartProd0,
			CheckSKUReq: 123,
			WantError: nil,
			WantRes:   0,
		},
	}
	ctrl := minimock.NewController(t)
	repositoryMock := mock.NewRepositoryMock(ctrl)
	repositoryMock.CheckSKUMock.When(123).Then(CheckSKUResp1, nil)
	repositoryMock.CheckSKUMock.When(124).Then(CheckSKUResp2, nil)
	NewHandler := New(repositoryMock)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			repositoryMock.GetCartMock.Expect(tt.CartId).Return(tt.CartProd, nil)

			Cart, err := NewHandler.GetCart(tt.CartId)

			assert.Equal(t, tt.WantError, err)
			assert.Equal(t, tt.WantRes, Cart.TotalPrice)
		})
	}
}
