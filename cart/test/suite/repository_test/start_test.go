package repository_test

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/config"
	"route256/cart/internal/repository"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteRepo struct {
	suite.Suite
	Repo *repository.Repository
}

func (s *SuiteRepo) SetupSuite() {
	s.Repo = &repository.Repository{
		Carts:  make(repository.Carts),
		Config: config.New(),
	}
}

func (s *SuiteRepo) SetupTest() {
	products := []struct {
		UserId int64
		SKU    int64
		Count  uint16
	}{
		{
			UserId: 12,
			SKU:    1076963,
			Count:  1,
		},
		{
			UserId: 12,
			SKU:    773297411,
			Count:  2,
		},
		{
			UserId: 12,
			SKU:    99999999, // Отсутствует в хранилище
			Count:  3,
		},
	}

	for _, prod := range products {
		product := &model.Product{
			SKU:   prod.SKU,
			Count: prod.Count,
		}

		err := s.Repo.AddProductCart(context.Background(), product, prod.UserId)
		if err != nil {
			s.T().Error("SetupTest")
		}
	}
}

func (s *SuiteRepo) TearDownTest() {
	s.Repo = &repository.Repository{
		Carts:  make(repository.Carts),
		Config: config.New(),
	}

}

func TestSuiteAll(t *testing.T) {
	suite.Run(t, new(SuiteRepo))
}
