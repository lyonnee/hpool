package hpool

import "sync"

type HPool[T any] struct {
	pool *sync.Pool
}

func New[T any](newFn func() T) *HPool[T] {
	return &HPool[T]{
		pool: &sync.Pool{
			New: func() any {
				return newFn()
			},
		},
	}
}

func (p *HPool[T]) Get() T {
	return p.pool.Get().(T)
}

func (p *HPool[T]) Put(obj T) {
	p.pool.Put(obj)
}
