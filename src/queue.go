package queue

import (
	"errors"
)

// Queue представляет кольцевую очередь строк
type Queue struct {
	data     []string
	capacity int
	front    int
	rear     int
	size     int
}

// New создает новую очередь с указанной емкостью
func New(capacity int) *Queue {
	return &Queue{
		data:     make([]string, capacity),
		capacity: capacity,
		front:    0,
		rear:     -1,
		size:     0,
	}
}

// NewDefault создает новую очередь с емкостью по умолчанию (30)
func NewDefault() *Queue {
	return New(30)
}

// IsEmpty проверяет, пуста ли очередь
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Push добавляет элемент в конец очереди
func (q *Queue) Push(value string) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}
	
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = value
	q.size++
	
	return nil
}

// Pop удаляет и возвращает элемент из начала очереди
func (q *Queue) Pop() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("queue is empty")
	}
	
	value := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	
	return value, nil
}

// Peek возвращает элемент из начала очереди без удаления
func (q *Queue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("queue is empty")
	}
	
	return q.data[q.front], nil
}

// Size возвращает текущее количество элементов в очереди
func (q *Queue) Size() int {
	return q.size
}

// Capacity возвращает емкость очереди
func (q *Queue) Capacity() int {
	return q.capacity
}
