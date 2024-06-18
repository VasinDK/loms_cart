package repository_test

import (
	"context"
	"route256/cart/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *SuiteRepo) TestDeleteProductCart() {
	tests := []struct {
		Name         string
		SKU          int64
		CartId       int64
		WantCountRes uint16
		WantErr      error
	}{
		{
			Name:         "SKU есть, корзина есть",
			SKU:          1076963,
			CartId:       12,
			WantCountRes: 0,
			WantErr:      nil,
		},
		{
			Name:         "SKU нет, корзина есть",
			SKU:          8888888,
			CartId:       12,
			WantCountRes: 0,
			WantErr:      nil,
		},
		{
			Name:         "Корзины нет",
			SKU:          1076963,
			CartId:       1200,
			WantCountRes: 0,
			WantErr:      nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			err := s.Repo.DeleteProductCart(context.Background(), tt.CartId, tt.SKU)
			assert.Equal(t, err, tt.WantErr)

			getProd := &model.Product{
				SKU: tt.SKU,
			}
			item, err := s.Repo.GetProductCart(context.Background(), getProd, tt.CartId)
			if err != nil {
				t.Error("Ошибка s.Repo.GetProductCart внутри TestDeleteProductCart")
			}
			assert.Equal(t, item.Count, tt.WantCountRes)
		})
	}
}
