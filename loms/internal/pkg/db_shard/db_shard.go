package db_shard

import (
	"context"
	"fmt"
	"hash"
	"route256/loms/internal/model"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conf interface {
	GetDBConnect() *[]string
	GetSequenceShift() string
	GetMainShard() string
}

type ShardManager struct {
	shards        []*pgxpool.Pool
	sequenceShift int64 // id начинается не с 1 а с 1000. Тут указывается сдвиг последовательности
	mainShard     int
	mu            sync.RWMutex
	hasher        hash.Hash32
}

func New(ctx context.Context, config Conf) (*ShardManager, error) {
	shift, err := strconv.ParseInt(config.GetSequenceShift(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("config.GetSequenceShift %w", err)
	}

	mainShard, err := strconv.ParseInt(config.GetMainShard(), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("config.GetMainShard %w", err)
	}

	sm := &ShardManager{
		sequenceShift: shift,
		mainShard:     int(mainShard),
	}

	for i, connStr := range *config.GetDBConnect() {
		if connStr == "" {
			return nil, model.ErrStrConnIsEmpty
		}

		connection, err := NewConn(ctx, i, connStr)
		if err != nil {
			return nil, fmt.Errorf("NewConn %w", err)
		}

		sm.shards = append(sm.shards, connection)
	}

	sm.SetHasher()

	return sm, nil
}

func NewConn(ctx context.Context, i int, connStr string) (*pgxpool.Pool, error) {
	connection, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New %w", err)
	}

	err = connection.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("conn.Ping %w", err)
	}

	return connection, nil
}

func (sm *ShardManager) Close() {
	for i := range sm.shards {
		sm.shards[i].Close()
	}
}
