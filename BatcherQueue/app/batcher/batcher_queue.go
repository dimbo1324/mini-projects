package batcher

import (
	"time"
)

// Handler — функция, вызываемая при накоплении батча.
type Handler[T any] func([]T)

// Batcher накапливает элементы и отдает их пачкой.
type Batcher[T any] struct {
}

// NewBatcher создает новый Batcher.
//
//   - flushSize — максимальное количество элементов в батче,
//   - interval — максимальное время ожидания перед сбросом.
//   - handler - пользовательский обработчик пачки
func NewBatcher[T any](capacity int, interval time.Duration, handler Handler[T]) *Batcher[T] {
	return &Batcher[T]{}
}

// Add добавляет элемент в батчер.
func (b *Batcher[T]) Add(items ...T) {

}

// Close завершает батчер с сохранением оставшихся элементов.
// Ожидает завершения всех handler'ов.
func (b *Batcher[T]) Close() {

}
