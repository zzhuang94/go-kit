package algo

import "cmp"

// Sort 对切片进行排序（使用快速排序）/ Sort slice using quicksort
func Sort[T cmp.Ordered](slice []T) {
	if len(slice) < 2 {
		return
	}
	quickSort(slice, 0, len(slice)-1)
}

// SortFunc 使用自定义比较函数对切片进行排序 / Sort slice using custom comparison function
func SortFunc[T any](slice []T, less func(a, b T) bool) {
	if len(slice) < 2 {
		return
	}
	quickSortFunc(slice, 0, len(slice)-1, less)
}

// SortDesc 对切片进行降序排序 / Sort slice in descending order
func SortDesc[T cmp.Ordered](slice []T) {
	if len(slice) < 2 {
		return
	}
	quickSortDesc(slice, 0, len(slice)-1)
}

// IsSorted 检查切片是否已排序 / Check if slice is sorted
func IsSorted[T cmp.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

// IsSortedFunc 使用自定义比较函数检查切片是否已排序 / Check if slice is sorted using custom comparison function
func IsSortedFunc[T any](slice []T, less func(a, b T) bool) bool {
	for i := 1; i < len(slice); i++ {
		if less(slice[i], slice[i-1]) {
			return false
		}
	}
	return true
}

// 快速排序实现 / Quicksort implementation
func quickSort[T cmp.Ordered](slice []T, low, high int) {
	if low < high {
		pivot := partition(slice, low, high)
		quickSort(slice, low, pivot-1)
		quickSort(slice, pivot+1, high)
	}
}

func partition[T cmp.Ordered](slice []T, low, high int) int {
	pivot := slice[high]
	i := low - 1
	for j := low; j < high; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

// 快速排序实现（自定义比较函数）/ Quicksort implementation with custom comparison function
func quickSortFunc[T any](slice []T, low, high int, less func(a, b T) bool) {
	if low < high {
		pivot := partitionFunc(slice, low, high, less)
		quickSortFunc(slice, low, pivot-1, less)
		quickSortFunc(slice, pivot+1, high, less)
	}
}

func partitionFunc[T any](slice []T, low, high int, less func(a, b T) bool) int {
	pivot := slice[high]
	i := low - 1
	for j := low; j < high; j++ {
		if less(slice[j], pivot) {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

// 快速排序实现（降序）/ Quicksort implementation (descending)
func quickSortDesc[T cmp.Ordered](slice []T, low, high int) {
	if low < high {
		pivot := partitionDesc(slice, low, high)
		quickSortDesc(slice, low, pivot-1)
		quickSortDesc(slice, pivot+1, high)
	}
}

func partitionDesc[T cmp.Ordered](slice []T, low, high int) int {
	pivot := slice[high]
	i := low - 1
	for j := low; j < high; j++ {
		if slice[j] > pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}
