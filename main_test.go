package main

import (
	"fmt"
	"sync"
	"testing"

	"AlekseyPromet/algo/lockfree/internal/cache"
	"AlekseyPromet/algo/lockfree/internal/lfstack"
	"AlekseyPromet/algo/lockfree/internal/lrucache"
	"AlekseyPromet/algo/lockfree/internal/ringbuffer"
)

func BenchmarkNewLockFreeStack(b *testing.B) {

	nfs := lfstack.NewLockFreeStack()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nfs.Push(i)
		got, _ := nfs.Pop()
		if got != i {
			b.Fatalf("unexpected popped value: got %d, want %d", got, i)
		}
	}
}

func BenchmarkNewLockFreeStackManyGoroutines(b *testing.B) {

	nfs := lfstack.NewLockFreeStack()
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				nfs.Push(i)
			}
		}(i)
	}

	for i := 0; i < b.N; i++ {
		go func(i int) {
			wg.Done()
			for j := 0; j < 10; j++ {
				nfs.Pop()
			}
		}(i)
	}
	wg.Wait()
}

func BenchmarkNewLRUCache(b *testing.B) {

	lru := lrucache.NewLRUCache(128)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ii := fmt.Sprintf("%d", i)
		lru.Set(ii, i)
		lru.Get(ii)
	}
}

func BenchmarkNewLRUCacheFull(b *testing.B) {

	lru := lrucache.NewLRUCache(128)
	for i := 0; i < 128; i++ {
		ii := fmt.Sprintf("%d", i)
		lru.Set(ii, i)
	}

	b.ResetTimer()
	for i := 0; i < 128+b.N; i++ {
		ii := fmt.Sprintf("%d", i)
		lru.Set(ii, i)
		lru.Get(ii)
	}
}

func BenchmarkNewCache(b *testing.B) {
	c := cache.NewCache()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprintf("%d", i), i)
		c.Get(fmt.Sprintf("%d", i))
	}

}

func BenchmarkNewRingBuffer(b *testing.B) {
	rb := ringbuffer.NewRingBuffer(128)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rb.Push(i)
		rb.Pop()
	}
}

func BenchmarkNewRingBufferFull(b *testing.B) {
	rb := ringbuffer.NewRingBuffer(128)

	for i := 0; i < 128; i++ {
		rb.Push(i)
	}

	b.ResetTimer()
	for i := 128; i < b.N; i++ {
		rb.Push(i)
		rb.Pop()
	}
}

func TestRunLockFreeStack(t *testing.T) {
	RunLockFreeStack()
}

func TestRunLRUCache(t *testing.T) {
	RunLockFreeCache()
}

func BenchmarkMain(b *testing.B) {
	parallerism := 32
	b.Logf("parallerism = %d\n", parallerism)
	b.SetParallelism(parallerism)
	b.Run("[ Sync map   ]", BenchmarkNewCache)
	b.Run("[ LRU cache  ] full lock-free", BenchmarkNewLRUCacheFull)
	b.Run("[ LRU cache  ] lock-free", BenchmarkNewLRUCache)
	b.Run("[ Stack      ] lock-free", BenchmarkNewLockFreeStack)
	b.Run("[ Ring buffer] lock-free", BenchmarkNewRingBuffer)
}
