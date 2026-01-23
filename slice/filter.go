package slice

// Filter 过滤切片，保留满足条件的元素 / Filter slice, keep elements matching condition
func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Find 查找第一个满足条件的元素，不存在返回零值 / Find first element matching condition, return zero value if not found
func Find[T any](slice []T, fn func(T) bool) (T, bool) {
	var zero T
	for _, v := range slice {
		if fn(v) {
			return v, true
		}
	}
	return zero, false
}

// FindIndex 查找第一个满足条件的元素索引，不存在返回 -1 / Find first element index matching condition, return -1 if not found
func FindIndex[T any](slice []T, fn func(T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

// FindLast 从后往前查找第一个满足条件的元素，不存在返回零值 / Find last element matching condition, return zero value if not found
func FindLast[T any](slice []T, fn func(T) bool) (T, bool) {
	var zero T
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i]) {
			return slice[i], true
		}
	}
	return zero, false
}

// FindLastIndex 从后往前查找第一个满足条件的元素索引，不存在返回 -1 / Find last element index matching condition, return -1 if not found
func FindLastIndex[T any](slice []T, fn func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i]) {
			return i
		}
	}
	return -1
}

// Every 检查是否所有元素都满足条件 / Check if all elements match condition
func Every[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Some 检查是否至少有一个元素满足条件 / Check if at least one element matches condition
func Some[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// Count 统计满足条件的元素个数 / Count elements matching condition
func Count[T any](slice []T, fn func(T) bool) int {
	count := 0
	for _, v := range slice {
		if fn(v) {
			count++
		}
	}
	return count
}

// GroupBy 按指定函数分组 / Group by function
func GroupBy[T any, K comparable](slice []T, fn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
