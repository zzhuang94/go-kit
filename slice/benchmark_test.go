package slice

import (
	"testing"
)

// 专门针对 []string 的实现，用于性能对比
func containsString(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func indexOfString(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func uniqueString(slice []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// 生成测试数据
func generateStringSlice(size int) []string {
	result := make([]string, size)
	for i := 0; i < size; i++ {
		result[i] = string(rune('a' + i%26))
	}
	return result
}

// Benchmark Contains - 泛型版本
func BenchmarkContains_Generic(b *testing.B) {
	slice := generateStringSlice(1000)
	item := "z"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(slice, item)
	}
}

// Benchmark Contains - 专用 []string 版本
func BenchmarkContains_String(b *testing.B) {
	slice := generateStringSlice(1000)
	item := "z"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsString(slice, item)
	}
}

// Benchmark IndexOf - 泛型版本
func BenchmarkIndexOf_Generic(b *testing.B) {
	slice := generateStringSlice(1000)
	item := "z"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IndexOf(slice, item)
	}
}

// Benchmark IndexOf - 专用 []string 版本
func BenchmarkIndexOf_String(b *testing.B) {
	slice := generateStringSlice(1000)
	item := "z"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexOfString(slice, item)
	}
}

// Benchmark Unique - 泛型版本
func BenchmarkUnique_Generic(b *testing.B) {
	slice := generateStringSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unique(slice)
	}
}

// Benchmark Unique - 专用 []string 版本
func BenchmarkUnique_String(b *testing.B) {
	slice := generateStringSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uniqueString(slice)
	}
}

// Benchmark Filter - 泛型版本
func BenchmarkFilter_Generic(b *testing.B) {
	slice := generateStringSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(slice, func(s string) bool {
			return len(s) > 0
		})
	}
}

// Benchmark Filter - 手动内联版本（避免函数调用开销）
func BenchmarkFilter_Manual(b *testing.B) {
	slice := generateStringSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result []string
		for _, v := range slice {
			if len(v) > 0 {
				result = append(result, v)
			}
		}
		_ = result
	}
}
