package lfstack

import (
	"sync/atomic"
	"unsafe"
)

// Node представляет узел стека.
type Node struct {
	value interface{}
	next  unsafe.Pointer // Атомарный указатель на следующий узел
}

// LockFreeStack представляет lock-free стек.
type LockFreeStack struct {
	top unsafe.Pointer // Атомарный указатель на верхний элемент стека
}

// NewLockFreeStack создаёт новый lock-free стек.
func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

// Push добавляет элемент в стек.
func (s *LockFreeStack) Push(value interface{}) {
	newNode := &Node{value: value, next: nil}
	for {
		// Читаем текущую вершину стека
		top := atomic.LoadPointer(&s.top)
		// Устанавливаем следующий элемент нового узла как текущую вершину
		newNode.next = top
		// Пытаемся атомарно заменить вершину стека на новый узел
		if atomic.CompareAndSwapPointer(&s.top, top, unsafe.Pointer(newNode)) {
			break // Успешно добавили элемент
		}
		// Если CAS не удался, повторяем попытку
	}
}

// Pop извлекает элемент из стека.
func (s *LockFreeStack) Pop() (interface{}, bool) {
	for {
		// Читаем текущую вершину стека
		top := atomic.LoadPointer(&s.top)
		if top == nil {
			return nil, false // Стек пуст
		}
		// Читаем следующий элемент
		next := (*Node)(top).next
		// Пытаемся атомарно заменить вершину стека на следующий элемент
		if atomic.CompareAndSwapPointer(&s.top, top, next) {
			return (*Node)(top).value, true // Успешно извлекли элемент
		}
		// Если CAS не удался, повторяем попытку
	}
}
