package slice

// Map 映射转换切片 / Map transform slice
func Map[T, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Reduce 归约切片 / Reduce slice
func Reduce[T, R any](slice []T, initial R, fn func(R, T) R) R {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// FlatMap 扁平映射 / Flat map
func FlatMap[T, R any](slice []T, fn func(T) []R) []R {
	var result []R
	for _, v := range slice {
		result = append(result, fn(v)...)
	}
	return result
}

// Partition 将切片分为两部分：满足条件的和不满足条件的 / Partition slice into two parts: matching and non-matching
func Partition[T any](slice []T, fn func(T) bool) ([]T, []T) {
	var trueSlice, falseSlice []T
	for _, v := range slice {
		if fn(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return trueSlice, falseSlice
}
