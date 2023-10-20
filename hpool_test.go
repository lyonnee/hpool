package hpool

import (
	"sync"
	"testing"
)

// 定义测量使用的对象
type DataObject struct {
	name  string
	value int
}

// 测试标准sync.Pool
func BenchmarkStdPool(b *testing.B) {
	pool := &sync.Pool{
		New: func() interface{} {
			return &DataObject{}
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		object := pool.Get().(*DataObject)
		object.name = "lyon"
		object.value = 26
		pool.Put(object)
	}
}

// 测试自定义池
func BenchmarkHPool(b *testing.B) {
	pool := New(func() *DataObject {
		return &DataObject{}
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		object := pool.Get()
		object.name = "lyon"
		object.value = 26
		pool.Put(object)
	}
}
