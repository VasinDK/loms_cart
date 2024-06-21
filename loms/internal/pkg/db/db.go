package db

import (
	"context"
	"fmt"
	"route256/loms/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conf interface {
	GetDBConnect() string
}

func NewConn(config Conf) (*pgxpool.Pool, error) {
	if config.GetDBConnect() == "" {
		return nil, model.ErrStrConnIsEmpty
	}

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, config.GetDBConnect())
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New %w", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("conn.Ping %w", err)
	}

	return conn, nil
}
