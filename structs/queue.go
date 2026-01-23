package structs

// Queue 队列数据结构 / Queue data structure
type Queue[T any] struct {
	items []T
}

// NewQueue 创建新队列 / Create new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue 入队 / Enqueue item
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue 出队 / Dequeue item
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Peek 查看队首元素 / Peek at front element
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	return q.items[0], true
}

// Size 返回队列大小 / Return queue size
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// IsEmpty 检查队列是否为空 / Check if queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Clear 清空队列 / Clear queue
func (q *Queue[T]) Clear() {
	q.items = q.items[:0]
}
