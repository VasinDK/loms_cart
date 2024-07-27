package db_shard

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spaolacci/murmur3"
)

// SetHasher - созданет хешер
func (sm *ShardManager) SetHasher() {
	sm.hasher = murmur3.New32()
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
	sm.mu.Lock()
	defer sm.mu.Unlock()
	defer sm.hasher.Reset()

	_, err := sm.hasher.Write([]byte(key))
	if err != nil {
		return 0, fmt.Errorf("sm.hasher.Write %w", err)
	}

	return sm.hasher.Sum32() % uint32(len(sm.shards)), nil
}

// GetMainShard - возвращает главный шард
func (sm *ShardManager) GetMainShard() int {
	return sm.mainShard
}
