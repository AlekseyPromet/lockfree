package ringbuffer

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// RingBuffer представляет lock-free кольцевой буфер.
type RingBuffer struct {
	data     []unsafe.Pointer // Слайс для хранения элементов
	size     int64            // Размер буфера
	producer int64            // Индекс для производителя
	consumer int64            // Индекс для потребителя
}

// NewRingBuffer создаёт новый кольцевой буфер заданного размера.
func NewRingBuffer(size int64) *RingBuffer {
	return &RingBuffer{
		data:     make([]unsafe.Pointer, size),
		size:     size,
		producer: 0,
		consumer: 0,
	}
}

// Push добавляет элемент в буфер.
func (rb *RingBuffer) Push(value interface{}) bool {
	for {
		// Читаем текущий индекс производителя
		producer := atomic.LoadInt64(&rb.producer)
		// Вычисляем следующий индекс
		nextProducer := (producer + 1) % rb.size
		// Если буфер полон, возвращаем false
		if nextProducer == atomic.LoadInt64(&rb.consumer) {
			return false
		}
		// Пытаемся атомарно обновить индекс производителя
		if atomic.CompareAndSwapInt64(&rb.producer, producer, nextProducer) {
			// Сохраняем значение в буфер
			atomic.StorePointer(&rb.data[producer], unsafe.Pointer(&value))
			return true
		}
		// Если CAS не удался, повторяем попытку
	}
}

// Pop извлекает элемент из буфера.
func (rb *RingBuffer) Pop() (interface{}, bool) {
	for {
		// Читаем текущий индекс потребителя
		consumer := atomic.LoadInt64(&rb.consumer)
		// Если буфер пуст, возвращаем false
		if consumer == atomic.LoadInt64(&rb.producer) {
			return nil, false
		}
		// Вычисляем следующий индекс
		nextConsumer := (consumer + 1) % rb.size
		// Пытаемся атомарно обновить индекс потребителя
		if atomic.CompareAndSwapInt64(&rb.consumer, consumer, nextConsumer) {
			// Читаем значение из буфера
			value := atomic.LoadPointer(&rb.data[consumer])
			return *(*interface{})(value), true
		}
		// Если CAS не удался, повторяем попытку
	}
}

func RunRingBuffer() {
	rb := NewRingBuffer(5)

	// Производитель добавляет элементы в буфер
	for i := 0; i < 5; i++ {
		success := rb.Push(i)
		if success {
			fmt.Println("Pushed:", i)
		} else {
			fmt.Println("Failed to push:", i)
		}
	}

	// Потребитель извлекает элементы из буфера
	for i := 0; i < 5; i++ {
		value, success := rb.Pop()
		if success {
			fmt.Println("Popped:", value)
		} else {
			fmt.Println("Failed to pop")
		}
	}
}