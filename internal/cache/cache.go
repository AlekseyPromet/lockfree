package cache

import (
	"sync"
)

// Cache представляет собой потокобезопасный кэш.
type Cache struct {
	store sync.Map
}

// NewCache создаёт новый кэш.
func NewCache() *Cache {
	return &Cache{
		store: sync.Map{},
	}
}

// Set добавляет или обновляет значение в кэше.
func (c *Cache) Set(key string, value interface{}) {
	c.store.Store(key, value)
}

// Get возвращает значение по ключу, если оно существует в кэше.
func (c *Cache) Get(key string) (interface{}, bool) {
	return c.store.Load(key)
}

// Delete удаляет значение из кэша по ключу.
func (c *Cache) Delete(key string) {
	c.store.Delete(key)
}

// Clear очищает весь кэш.
func (c *Cache) Clear() {
	c.store.Range(func(key, value interface{}) bool {
		c.store.Delete(key)
		return true
	})
}

// Range выполняет функцию для каждого элемента кэша.
func (c *Cache) Range(f func(key, value interface{}) bool) {
	c.store.Range(f)
}
