package structs

// Stack 栈数据结构 / Stack data structure
type Stack[T any] struct {
	items []T
}

// NewStack 创建新栈 / Create new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push 入栈 / Push item onto stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop 出栈 / Pop item from stack
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek 查看栈顶元素 / Peek at top element
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Size 返回栈大小 / Return stack size
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// IsEmpty 检查栈是否为空 / Check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear 清空栈 / Clear stack
func (s *Stack[T]) Clear() {
	s.items = s.items[:0]
}
