package add_product

import (
	"context"
	"fmt"
	"route256/cart/internal/model"
	"route256/cart/internal/service/item/add_product/mock"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	t.Parallel()
	product := &model.Product{
		SKU:   123,
		Name:  "Чай",
		Price: 13,
		Count: 31,
	}
	productCount0 := &model.Product{
		SKU:   123,
		Name:  "Чай",
		Price: 13,
		Count: 0,
	}
	tests := []struct {
		Name               string
		UserId             int64
		CheckSKUReq        int64
		CheckSKURespParam1 *model.Product
		CheckSKURespParam2 error
		GetProdReq         *model.Product
		GetProdRespParam1  *model.Product
		GetProdRespParam2  error
		AddProdReq         *model.Product
		AddProdResp        error
		WantError          error
	}{
		{
			Name:               "sku отсутствует в хранилище",
			UserId:             22,
			CheckSKUReq:        123,
			CheckSKURespParam1: &model.Product{},
			CheckSKURespParam2: model.ErrNoProductInStock,
			GetProdReq:         product,
			GetProdRespParam1:  &model.Product{},
			GetProdRespParam2:  nil,
			AddProdReq:         product,
			AddProdResp:        nil,
			WantError:          fmt.Errorf("s.Repository.CheckSKU %w", model.ErrNoProductInStock),
		},
		{
			Name:               "sku есть, ранее не добавлены в корзину",
			UserId:             22,
			CheckSKUReq:        123,
			CheckSKURespParam1: &model.Product{Price: 12345},
			CheckSKURespParam2: nil,
			GetProdReq:         product,
			GetProdRespParam1:  &model.Product{},
			GetProdRespParam2:  nil,
			AddProdReq:         product,
			AddProdResp:        nil,
			WantError:          nil,
		},
		{
			Name:               "sku есть, ранее добавлены в корзину",
			UserId:             22,
			CheckSKUReq:        123,
			CheckSKURespParam1: &model.Product{Price: 12345},
			CheckSKURespParam2: nil,
			GetProdReq:         product,
			GetProdRespParam1:  product,
			GetProdRespParam2:  nil,
			AddProdReq:         product,
			AddProdResp:        nil,
			WantError:          nil,
		},
		{
			Name:               "sku есть, count 0",
			UserId:             0,
			CheckSKUReq:        123,
			CheckSKURespParam1: &model.Product{Price: 12345},
			CheckSKURespParam2: nil,
			GetProdReq:         product,
			GetProdRespParam1:  product,
			GetProdRespParam2:  nil,
			AddProdReq:         productCount0,
			AddProdResp:        nil,
			WantError:          fmt.Errorf("AddProduct %w", fmt.Errorf("количество меньше 1")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := minimock.NewController(t)
			repoMock := mock.NewRepositoryMock(ctrl)
			repoMock.CheckSKUMock.Optional().Set(func(ctx context.Context, ch1 chan<- *model.Product, userId int64) error {
				ch1 <- tt.CheckSKURespParam1
				return tt.CheckSKURespParam2
			})

			repoMock.GetProductCartMock.Optional().Expect(context.Background(), tt.GetProdReq, tt.UserId).Return(tt.GetProdRespParam1, tt.GetProdRespParam2)
			repoMock.AddProductCartMock.Optional().Expect(context.Background(), tt.AddProdReq, tt.UserId).Return(tt.AddProdResp)
			repoMock.StockInfoMock.Optional().Expect(context.Background(), tt.CheckSKUReq).Return(1000, nil)

			NewRepo := New(repoMock)
			err := NewRepo.AddProduct(context.Background(), tt.AddProdReq, tt.UserId)
			assert.Equal(t, tt.WantError, err)
		})
	}
}
