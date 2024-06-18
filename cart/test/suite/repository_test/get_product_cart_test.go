package repository_test

import (
	"context"
	"route256/cart/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *SuiteRepo) TestGetProductCart() {
	tests := []struct {
		Name      string
		SKU       int64
		UserId    int64
		WantCount uint16
		WantErr   error
	}{
		{
			Name:      "Продукт есть, корзина есть",
			SKU:       1076963, // 1076963, 773297411
			UserId:    12,
			WantErr:   nil,
			WantCount: 1,
		},
		{
			Name:      "Продукта нет, корзина есть",
			SKU:       9999900000,
			UserId:    12,
			WantErr:   nil,
			WantCount: 0,
		},
		{
			Name:      "Продукта нет, корзины нет",
			SKU:       99999999,
			UserId:    12000,
			WantErr:   nil,
			WantCount: 0,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			product := &model.Product{
				SKU: tt.SKU,
			}

			item, err := s.Repo.GetProductCart(context.Background(), product, tt.UserId)

			assert.Equal(t, err, tt.WantErr)
			assert.Equal(t, item.Count, tt.WantCount)
		})
	}
}
