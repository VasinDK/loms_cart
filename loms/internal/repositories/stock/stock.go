package stock

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockItem struct {
	Sku        uint32 `json:"sku"`
	TotalCount uint64 `json:"total_count"`
	Reserved   uint64 `json:"reserved"`
}

type StockRepository struct {
	Conn *pgxpool.Pool
}

// New - создает новый репозиторий для StockRepository
func New(conn *pgxpool.Pool) *StockRepository {
	return &StockRepository{
		Conn: conn,
	}
}
