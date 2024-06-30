package repository_test

import (
	"context"
	"route256/cart/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *SuiteRepo) TestCheckSKU() {
	tests := []struct {
		Name    string
		SKU     int64
		WantRes *model.Product
		WantErr error
	}{
		{
			Name: "Существующий SKU",
			SKU:  773297411,
			WantRes: &model.Product{
				Name:  "Кроссовки Nike JORDAN",
				Price: 2202,
				SKU:   773297411,
			},
			WantErr: nil,
		},
		{
			Name:    "Не существующий SKU",
			SKU:     9999999,
			WantRes: nil,
			WantErr: model.ErrNoProductInStock,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			res, err := s.Repo.CheckSKU(context.Background(), tt.SKU)

			assert.Equal(t, res, tt.WantRes)
			assert.Equal(t, err, tt.WantErr)
		})
	}
}
