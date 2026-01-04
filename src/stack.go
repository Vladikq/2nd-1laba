package stack

import (
	"errors"
)

// Stack представляет стек строк
type Stack struct {
	data     []string
	capacity int
	head     int
}

// New создает новый стек с указанной емкостью
func New(capacity int) *Stack {
	return &Stack{
		data:     make([]string, capacity),
		capacity: capacity,
		head:     -1,
	}
}

// NewDefault создает новый стек с емкостью по умолчанию (30)
func NewDefault() *Stack {
	return New(30)
}

// Push добавляет элемент на вершину стека
func (s *Stack) Push(value string) error {
	if s.head == s.capacity-1 {
		return errors.New("stack is full")
	}
	
	s.head++
	s.data[s.head] = value
	
	return nil
}

// Pop удаляет и возвращает элемент с вершины стека
func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}
	
	value := s.data[s.head]
	s.head--
	
	return value, nil
}

// Peek возвращает элемент с вершины стека без удаления
func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("stack is empty")
	}
	
	return s.data[s.head], nil
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return s.head == -1
}

// Size возвращает текущее количество элементов в стеке
func (s *Stack) Size() int {
	return s.head + 1
}

// Capacity возвращает емкость стека
func (s *Stack) Capacity() int {
	return s.capacity
}

// Clear очищает стек
func (s *Stack) Clear() {
	s.head = -1
}
