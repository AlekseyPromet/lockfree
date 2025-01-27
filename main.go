package main

import (
	"fmt"

	"github.com/AlekseyPromet/lockfree/internal/cache"
	"github.com/AlekseyPromet/lockfree/internal/lfqueue"
	"github.com/AlekseyPromet/lockfree/internal/lfstack"
	"github.com/AlekseyPromet/lockfree/internal/lrucache"
)

func RunLockFreeStack() {
	stack := lfstack.NewLockFreeStack()

	// Добавляем элементы в стек
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	// Извлекаем элементы из стека
	for i := 0; i < 4; i++ {
		val, ok := stack.Pop()
		if ok {
			fmt.Println("Popped:", val)
		} else {
			fmt.Println("Stack is empty")
		}
	}
}

func RunLockFreeCache() {
	cache := lrucache.NewLRUCache(2)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	if val, ok := cache.Get("key1"); ok {
		println("key1:", val.(string))
	}

	cache.Set("key3", "value3") // key2 будет удалён, так как кэш переполнен

	if _, ok := cache.Get("key2"); !ok {
		println("key2 not found")
	}
}

func RunCache() {
	cache := cache.NewCache()

	// Добавляем значения в кэш
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	// Получаем значения из кэша
	if val, ok := cache.Get("key1"); ok {
		fmt.Println("key1:", val)
	}

	// Удаляем значение из кэша
	cache.Delete("key2")

	// Проверяем, удалено ли значение
	if _, ok := cache.Get("key2"); !ok {
		fmt.Println("key2 not found")
	}

	// Очищаем кэш
	cache.Clear()

	// Проверяем, очищен ли кэш
	cache.Set("key3", "value3")
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	})
}

func RunLockFreeQueue() {
	q := lfqueue.NewLockFreeQueue()

	// Добавляем элементы в очередь
	q.Push(1)
	q.Push(2)
	q.Push(3)

	// Извлекаем элементы из очереди
	for i := 0; i < 4; i++ {
		val, ok := q.Pop()
		if ok {
			println(val.(int))
		} else {
			println("Queue is empty")
		}
	}
}

func main() {
	RunLockFreeStack()
	RunLockFreeCache()
	RunCache()
}
