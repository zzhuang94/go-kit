package structs

// Set 集合数据结构 / Set data structure
type Set[T comparable] struct {
	items map[T]bool
}

// NewSet 创建新集合 / Create new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]bool),
	}
}

// Add 添加元素 / Add element
func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

// Remove 删除元素 / Remove element
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Contains 检查是否包含元素 / Check if contains element
func (s *Set[T]) Contains(item T) bool {
	return s.items[item]
}

// Size 返回集合大小 / Return set size
func (s *Set[T]) Size() int {
	return len(s.items)
}

// IsEmpty 检查集合是否为空 / Check if set is empty
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear 清空集合 / Clear set
func (s *Set[T]) Clear() {
	s.items = make(map[T]bool)
}

// ToSlice 转换为切片 / Convert to slice
func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

// Union 求并集 / Get union
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		result.Add(item)
	}
	for item := range other.items {
		result.Add(item)
	}
	return result
}

// Intersect 求交集 / Get intersection
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Diff 求差集（s 中有但 other 中没有的元素）/ Get difference (elements in s but not in other)
func (s *Set[T]) Diff(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
