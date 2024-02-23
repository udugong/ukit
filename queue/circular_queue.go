package queue

import "errors"

// CircularQueue 循环队列.
type CircularQueue[T any] struct {
	capacity uint // 容量
	head     uint // 指向队头的索引
	tail     uint // 指向队尾的索引
	data     []T  // 队列中的元素
}

// NewCircularQueue 创建一个循环队列.
// capacity 必须大于0 否则会 panic.
// 因为 tail 指向的位置实际上是没有数据的
// 所以 data 的实际容量为 capacity+1.
func NewCircularQueue[T any](capacity uint) *CircularQueue[T] {
	if capacity < 1 {
		panic("ukit: 队列容量必须为正数")
	}
	realCap := capacity + 1
	return &CircularQueue[T]{
		capacity: realCap,
		data:     make([]T, realCap),
	}
}

var ErrFullQueue = errors.New("ukit: 队列已满")

// Enqueue 入队.
// 如果队列已满则返回 ErrFullQueue 错误.
func (c *CircularQueue[T]) Enqueue(val T) error {
	if c.IsFull() {
		return ErrFullQueue
	}
	c.data[c.tail] = val
	c.tail = (c.tail + 1) % c.capacity
	return nil
}

var ErrEmptyQueue = errors.New("ukit: 队列为空")

// Dequeue 出队.
// 如果队列为空则返回 ErrEmptyQueue 错误.
func (c *CircularQueue[T]) Dequeue() (T, error) {
	var val T
	if c.IsEmpty() {
		return val, ErrEmptyQueue
	}
	val = c.data[c.head]
	c.head = (c.head + 1) % c.capacity
	return val, nil
}

// Peek 查看队头元素.
// 如果队列为空则返回 ErrEmptyQueue 错误.
func (c *CircularQueue[T]) Peek() (T, error) {
	var t T
	if c.IsEmpty() {
		return t, ErrEmptyQueue
	}
	return c.data[c.head], nil
}

// IsEmpty 队列为空.
func (c *CircularQueue[T]) IsEmpty() bool {
	if c.head != c.tail {
		return false
	}
	return true
}

// IsFull 队列已满.
func (c *CircularQueue[T]) IsFull() bool {
	if (c.tail+1)%c.capacity == c.head {
		return true
	}
	return false
}
