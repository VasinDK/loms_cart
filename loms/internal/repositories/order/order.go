package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	Conn *pgxpool.Pool
}

// New - создает OrderRepository
func New(conn *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		Conn: conn,
	}
}
