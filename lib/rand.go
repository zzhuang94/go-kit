package lib

import (
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Choice 从切片中随机选择一个元素 / Randomly select an element from slice
func Choice[T any](slice []T) (T, bool) {
	var zero T
	if len(slice) == 0 {
		return zero, false
	}
	return slice[r.Intn(len(slice))], true
}

// Shuffle 随机打乱切片 / Randomly shuffle slice
func Shuffle[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	r.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
