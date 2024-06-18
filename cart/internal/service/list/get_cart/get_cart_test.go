package get_cart

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/internal/service/list/get_cart/mock"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetCart(t *testing.T) {
	t.Parallel()
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

	CheckSKUResp1 := &model.Product{ // SKU:   123
		Name:  "Чай",
		Price: 13,
	}

	CheckSKUResp2 := &model.Product{ // SKU:   124
		Name:  "Кофе",
		Price: 130,
	}
	CheckSKUResp0 := &model.Product{}

	tests := []struct {
		Name        string
		CartId      int64
		CartProd    map[int64]*model.Product
		CheckSKUReq int64
		WantError   error
		WantRes     uint32
	}{
		{
			Name:        "Типичное значение",
			CartId:      12,
			CartProd:    CartProd,
			CheckSKUReq: 123,
			WantError:   nil,
			WantRes:     156,
		},
		{
			Name:        "Пустая корзина",
			CartId:      2,
			CartProd:    map[int64]*model.Product{},
			CheckSKUReq: 123,
			WantError:   nil,
			WantRes:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := minimock.NewController(t)
			repositoryMock := mock.NewRepositoryMock(ctrl)
			repositoryMock.GetCartMock.Optional().Expect(context.Background(), tt.CartId).Return(tt.CartProd, nil)
			repositoryMock.CheckSKUMock.Optional().Set(func(ctx context.Context, i int64) (*model.Product, error) {
				switch i {
				case 123:
					return CheckSKUResp1, nil
				case 124:
					return CheckSKUResp2, nil
				default:
					return CheckSKUResp0, nil
				}
			})

			NewHandler := New(repositoryMock)
			Cart, err := NewHandler.GetCart(context.Background(), tt.CartId)

			assert.Equal(t, tt.WantError, err)
			assert.Equal(t, tt.WantRes, Cart.TotalPrice)
		})
	}
}
