package repository_test

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/pkg/errgroup_my"
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
			WantRes: &model.Product{},
			WantErr: model.ErrNoProductInStock,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			ch1 := make(chan *model.Product)

			eg, ctx := errgroup_my.WithContext(context.Background())
			eg.Go(func() error {
				return s.Repo.CheckSKU(ctx, ch1, tt.SKU)
			})

			res := <-ch1
			err := eg.Wait()

			close(ch1)

			assert.Equal(t, res, tt.WantRes)
			assert.Equal(t, err, tt.WantErr)
		})
	}
}
