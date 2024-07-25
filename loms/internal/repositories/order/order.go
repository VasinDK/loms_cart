package order

import (
	"route256/loms/internal/repositories/stock"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	Sm ShardManager
}

type ShardManager interface {
	GetShardIndexFromID(id int64) int
	GetShardIndex(key string) (uint32, error)
	Pick(index int) (*pgxpool.Pool, error)
	GetMainShard() int
}

var (
	RequestDBTotal              = stock.RequestDBTotal
	RequestTimeStatusCategoryBD = stock.RequestTimeStatusCategoryBD
)

// New - создает OrderRepository
func New(sm ShardManager) *OrderRepository {
	return &OrderRepository{
		Sm: sm,
	}
}
