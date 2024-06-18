package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *SuiteRepo) TestClearCart() {
	tests := []struct {
		Name        string
		CartIdDel   int64
		CartIdCheck int64
		WantCount   int
		WantErr     error
	}{
		{
			Name:        "Корзина",
			CartIdDel:   12,
			CartIdCheck: 12,
			WantCount:   1,
			WantErr:     nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			err := s.Repo.ClearCart(context.Background(), tt.CartIdDel)
			assert.Equal(t, err, tt.WantErr)

			items, err := s.Repo.GetCart(context.Background(), tt.CartIdCheck)
			if err != nil {
				t.Error("Ошибка s.Repo.GetCart внутри TestClearCart")
			}
			assert.Greater(t, tt.WantCount, len(items))
		})
	}
}
