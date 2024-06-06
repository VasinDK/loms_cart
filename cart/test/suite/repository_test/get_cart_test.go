package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *SuiteRepo) TestGetCart() {
	tests := []struct {
		Name      string
		CartId    int64
		WantCount int
		WantErr   error
	}{
		{
			Name:      "Корзина",
			CartId:    12,
			WantCount: 0,
			WantErr:   nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			items, err := s.Repo.GetCart(tt.CartId)
			assert.Equal(t, err, tt.WantErr)

			assert.Greater(t, len(items), tt.WantCount)
		})
	}
}
