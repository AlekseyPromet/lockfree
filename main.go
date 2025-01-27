package main

import (
	"AlekseyPromet/algo/lockfree/internal/cache"
	"AlekseyPromet/algo/lockfree/internal/lfstack"
	"AlekseyPromet/algo/lockfree/internal/lrucache"
	"fmt"
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

func main() {
	RunLockFreeStack()
	RunLockFreeCache()
	RunCache()
}
