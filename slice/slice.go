package slice

import (
	"math/rand"
	"time"
)

// Contains 检查切片是否包含指定元素 / Check if slice contains element
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// IndexOf 查找元素在切片中的索引，不存在返回 -1 / Find element index in slice, return -1 if not found
func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// LastIndexOf 从后往前查找元素在切片中的索引，不存在返回 -1 / Find element index from end, return -1 if not found
func LastIndexOf[T comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// Unique 去除重复元素 / Remove duplicate elements
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Reverse 反转切片 / Reverse slice
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		result[i] = slice[j]
	}
	return result
}

// Shuffle 随机打乱切片 / Randomly shuffle slice
func Shuffle[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

// Chunk 将切片分块 / Split slice into chunks
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Flatten 扁平化嵌套切片 / Flatten nested slices
func Flatten[T any](slices [][]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// Intersect 求两个切片的交集 / Get intersection of two slices
func Intersect[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	for _, v := range slice2 {
		set[v] = true
	}
	var result []T
	seen := make(map[T]bool)
	for _, v := range slice1 {
		if set[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Union 求两个切片的并集 / Get union of two slices
func Union[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	var result []T
	for _, v := range slice1 {
		if !set[v] {
			set[v] = true
			result = append(result, v)
		}
	}
	for _, v := range slice2 {
		if !set[v] {
			set[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Diff 求两个切片的差集（slice1 中有但 slice2 中没有的元素）/ Get difference of two slices (elements in slice1 but not in slice2)
func Diff[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	for _, v := range slice2 {
		set[v] = true
	}
	var result []T
	seen := make(map[T]bool)
	for _, v := range slice1 {
		if !set[v] && !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Remove 删除切片中所有匹配的元素 / Remove all matching elements
func Remove[T comparable](slice []T, item T) []T {
	var result []T
	for _, v := range slice {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

// RemoveAt 删除指定索引的元素 / Remove element at index
func RemoveAt[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

// Insert 在指定索引插入元素 / Insert element at index
func Insert[T any](slice []T, index int, item T) []T {
	if index < 0 {
		index = 0
	}
	if index > len(slice) {
		index = len(slice)
	}
	result := make([]T, 0, len(slice)+1)
	result = append(result, slice[:index]...)
	result = append(result, item)
	result = append(result, slice[index:]...)
	return result
}

// First 获取第一个元素，如果切片为空返回零值 / Get first element, return zero value if slice is empty
func First[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[0]
}

// Last 获取最后一个元素，如果切片为空返回零值 / Get last element, return zero value if slice is empty
func Last[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[len(slice)-1]
}

// Take 取前N个元素 / Take first N elements
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if n >= len(slice) {
		return slice
	}
	result := make([]T, n)
	copy(result, slice[:n])
	return result
}

// Drop 跳过前N个元素 / Drop first N elements
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return nil
	}
	return slice[n:]
}
