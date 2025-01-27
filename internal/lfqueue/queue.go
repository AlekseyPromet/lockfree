package lfqueue

import (
	"sync/atomic"
	"unsafe"
)

// Node представляет узел очереди
type Node struct {
	value interface{}
	next  unsafe.Pointer // Указатель на следующий узел (atomic)
}

// LockFreeQueue представляет lock-free очередь
type LockFreeQueue struct {
	head unsafe.Pointer // Указатель на голову очереди (atomic)
	tail unsafe.Pointer // Указатель на хвост очереди (atomic)
}

// NewLockFreeQueue создаёт новую lock-free очередь
func NewLockFreeQueue() *LockFreeQueue {
	dummy := &Node{} // Фиктивный узел
	return &LockFreeQueue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

// Enqueue добавляет элемент в очередь
func (q *LockFreeQueue) Enqueue(value interface{}) {
	newNode := &Node{value: value}

	for {
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*Node)(tail).next)

		// Если tail не изменился
		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				// Пытаемся добавить новый узел
				if atomic.CompareAndSwapPointer(&(*Node)(tail).next, next, unsafe.Pointer(newNode)) {
					// Перемещаем tail на новый узел
					atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newNode))
					return
				}
			} else {
				// Помогаем переместить tail, если он отстал
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			}
		}
	}
}

// Dequeue извлекает элемент из очереди
func (q *LockFreeQueue) Dequeue() (interface{}, bool) {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*Node)(head).next)

		// Если head не изменился
		if head == atomic.LoadPointer(&q.head) {
			// Если очередь пуста
			if head == tail {
				if next == nil {
					return nil, false
				}
				// Помогаем переместить tail, если он отстал
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			} else {
				// Читаем значение
				value := (*Node)(next).value

				// Пытаемся переместить head на следующий узел
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					return value, true
				}
			}
		}
	}
}
