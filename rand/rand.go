package rand

import (
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Int 生成随机整数 / Generate random integer
func Int() int {
	return r.Int()
}

// IntRange 生成指定范围的随机整数 [min, max) / Generate random integer in range [min, max)
func IntRange(min, max int) int {
	if min >= max {
		return min
	}
	return min + r.Intn(max-min)
}

// Float64 生成随机浮点数 [0.0, 1.0) / Generate random float [0.0, 1.0)
func Float64() float64 {
	return r.Float64()
}

// String 生成随机字符串（字母数字）/ Generate random string (alphanumeric)
func String(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return StringWithCharset(length, charset)
}

// StringWithCharset 使用指定字符集生成随机字符串 / Generate random string with specified charset
func StringWithCharset(length int, charset string) string {
	if length <= 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

// Bytes 生成随机字节数组 / Generate random byte array
func Bytes(length int) []byte {
	if length <= 0 {
		return nil
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(r.Intn(256))
	}
	return b
}

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
