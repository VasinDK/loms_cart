package ring_cache

import (
	"context"
	"sync"
)

type Cacher struct {
	SizeBuffer int64
	Cache      sync.Map
	RingBuffer []string
	mx         sync.RWMutex
}

func (cr *Cacher) Set(ctx context.Context, key string, value any) {
	if int64(len(cr.RingBuffer)) >= cr.SizeBuffer {
		cr.mx.Lock()
		oldKey := cr.RingBuffer[cr.SizeBuffer-1]
		cr.RingBuffer = append(cr.RingBuffer, key)
		cr.RingBuffer = cr.RingBuffer[1:]
		cr.mx.Unlock()

		cr.Cache.Delete(oldKey)
		cr.Cache.Store(key, value)
		return
	}

	cr.mx.Lock()
	cr.RingBuffer = append(cr.RingBuffer, key)
	cr.mx.Unlock()

	cr.Cache.Store(key, value)
}

func (cr *Cacher) Get(ctx context.Context, key string) (any, bool) {
	res, ok := cr.Cache.Load(key)
	return res, ok
}
