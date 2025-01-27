package lrucache

import (
	"container/list"
	"errors"
	"sync"
	"sync/atomic"
	"unsafe"
)

// LRUCache представляет собой lock-free LRU-кэш.
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mu       sync.Mutex
}

// entry представляет элемент кэша.
type entry struct {
	key   string
	value unsafe.Pointer // Атомарный указатель на значение
}

// NewLRUCache создаёт новый LRU-кэш с заданной ёмкостью.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// Get возвращает значение по ключу, если оно существует в кэше.
func (l *LRUCache) Get(key string) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		value := atomic.LoadPointer(&elem.Value.(*entry).value)
		return *(*interface{})(value), true
	}
	return nil, false
}

// Set добавляет или обновляет значение в кэше.
func (l *LRUCache) Set(key string, value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if elem, ok := l.cache[key]; ok {
		// Обновляем значение атомарно
		atomic.StorePointer(&elem.Value.(*entry).value, unsafe.Pointer(&value))
		l.list.MoveToFront(elem)
	} else {
		// Если кэш заполнен, удаляем последний элемент
		if l.list.Len() >= l.capacity {
			last := l.list.Back()
			if last != nil {
				delete(l.cache, last.Value.(*entry).key)
				l.list.Remove(last)
			}
		}
		// Добавляем новый элемент
		newEntry := &entry{key: key, value: unsafe.Pointer(&value)}
		elem := l.list.PushFront(newEntry)
		l.cache[key] = elem
	}
}

var ErrKeyNotFound = errors.New("key not found")

func (l *LRUCache) Pop(key string) (*entry, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if elem, ok := l.cache[key]; ok {
		value := atomic.LoadPointer(&elem.Value.(*entry).value)
		delete(l.cache, key)
		l.list.Remove(elem)
		return &entry{key: key, value: value}, nil
	}
	return nil, ErrKeyNotFound
}
