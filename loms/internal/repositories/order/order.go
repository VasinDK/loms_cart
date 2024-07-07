package order

import (
	"route256/loms/internal/repositories/stock"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	Conn *pgxpool.Pool
}

var (
	RequestDBTotal              = stock.RequestDBTotal
	RequestTimeStatusCategoryBD = stock.RequestTimeStatusCategoryBD
)

// New - создает OrderRepository
func New(conn *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		Conn: conn,
	}
}
