package db_shard

import "fmt"

var (
	ErrShardIndexOutOfRange = fmt.Errorf("shard index is out of range")
	ErrGettingIdShift       = fmt.Errorf("error getting ID shift")
)
