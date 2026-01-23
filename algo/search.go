package algo

import "cmp"

// BinarySearch 二分查找，返回目标值的索引，如果不存在返回 -1 / Binary search, return index of target, -1 if not found
func BinarySearch[T cmp.Ordered](slice []T, target T) int {
	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// BinarySearchFunc 使用自定义比较函数进行二分查找 / Binary search with custom comparison function
func BinarySearchFunc[T any](slice []T, target T, cmp func(a, b T) int) int {
	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		compare := cmp(slice[mid], target)
		if compare == 0 {
			return mid
		} else if compare < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// LinearSearch 线性查找，返回目标值的索引，如果不存在返回 -1 / Linear search, return index of target, -1 if not found
func LinearSearch[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// FindFirst 查找第一个满足条件的元素索引，不存在返回 -1 / Find first element index matching condition, return -1 if not found
func FindFirst[T any](slice []T, fn func(T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

// FindLast 查找最后一个满足条件的元素索引，不存在返回 -1 / Find last element index matching condition, return -1 if not found
func FindLast[T any](slice []T, fn func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i]) {
			return i
		}
	}
	return -1
}

// LowerBound 查找第一个大于等于目标值的元素索引 / Find first element index >= target
func LowerBound[T cmp.Ordered](slice []T, target T) int {
	left, right := 0, len(slice)
	for left < right {
		mid := left + (right-left)/2
		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// UpperBound 查找第一个大于目标值的元素索引 / Find first element index > target
func UpperBound[T cmp.Ordered](slice []T, target T) int {
	left, right := 0, len(slice)
	for left < right {
		mid := left + (right-left)/2
		if slice[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
