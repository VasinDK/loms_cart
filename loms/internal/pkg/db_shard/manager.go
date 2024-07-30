package db_shard

import (
	"fmt"
	"hash"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spaolacci/murmur3"
)

// SetHasher - созданет хешер
func (sm *ShardManager) SetHasher() {
	sm.hasher = sync.Pool{
		New: func() any {
			return murmur3.New32()
		},
	}
}

// Pick - возвращает полу соединений конкретного шарда
func (sm *ShardManager) Pick(index int) (*pgxpool.Pool, error) {
	if index < len(sm.shards) {
		return sm.shards[index], nil
	}

	return nil, fmt.Errorf("%w: given index=%d, len=%d", ErrShardIndexOutOfRange, index, len(sm.shards))
}

// GetShardIndexFromID - получает номер шарда по id элемента шарда
func (sm *ShardManager) GetShardIndexFromID(id int64) int {
	res := int(id % sm.sequenceShift)
	return res
}

// GetShardIndex - получает номер шарда по ключу
func (sm *ShardManager) GetShardIndex(key string) (uint32, error) {
	hash := sm.hasher.Get().(hash.Hash32)
	defer func() {
		hash.Reset()
		sm.hasher.Put(hash)
	}()

	_, err := hash.Write([]byte(key))
	if err != nil {
		return 0, fmt.Errorf("hash.Write %w", err)
	}

	res := hash.Sum32() % uint32(len(sm.shards))

	return res, nil
}

// GetMainShard - возвращает главный шард
func (sm *ShardManager) GetMainShard() int {
	return sm.mainShard
}
