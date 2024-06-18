package repository_test

import (
	"context"
	"route256/cart/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *SuiteRepo) TestAddProductCart() {
	tests := []struct {
		Name         string
		SKU          int64
		Count        uint16
		UserId       int64
		WantCountRes uint16
		WantErr      error
	}{
		{
			Name:         "Продукт ранее не добавлялся",
			SKU:          773297411, // 1076963, 773297411
			Count:        3,
			UserId:       11,
			WantErr:      nil,
			WantCountRes: 3,
		},
		{
			Name:         "Продукт ранее добавлялся",
			SKU:          773297411,
			Count:        8,
			UserId:       11,
			WantErr:      nil,
			WantCountRes: 8,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			product := &model.Product{
				SKU:   tt.SKU,
				Count: tt.Count,
			}

			err := s.Repo.AddProductCart(context.Background(), product, tt.UserId)
			assert.Equal(t, err, tt.WantErr)

			item, err := s.Repo.GetProductCart(context.Background(), product, tt.UserId)
			if err != nil {
				t.Error("Ошибка s.Repo.GetProductCart внутри TestAddProductCart")
			}

			assert.Equal(t, item.SKU, tt.SKU)
			assert.Equal(t, item.Count, tt.WantCountRes)
		})
	}
}
