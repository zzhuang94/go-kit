package algo

import (
	"testing"
)

// 专门针对 string 的 Set 实现，用于性能对比
type StringSet struct {
	items map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{
		items: make(map[string]bool),
	}
}

func (s *StringSet) Add(item string) {
	s.items[item] = true
}

func (s *StringSet) Contains(item string) bool {
	return s.items[item]
}

func (s *StringSet) Size() int {
	return len(s.items)
}

// 生成测试数据
func generateStringSet(size int) []string {
	result := make([]string, size)
	for i := 0; i < size; i++ {
		result[i] = string(rune('a' + i%26))
	}
	return result
}

// Benchmark Set Add - 泛型版本
func BenchmarkSetAdd_Generic(b *testing.B) {
	set := NewSet[string]()
	items := generateStringSet(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range items {
			set.Add(item)
		}
		set.Clear()
	}
}

// Benchmark Set Add - 专用 string 版本
func BenchmarkSetAdd_String(b *testing.B) {
	set := NewStringSet()
	items := generateStringSet(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range items {
			set.Add(item)
		}
		set.items = make(map[string]bool)
	}
}

// Benchmark Set Contains - 泛型版本
func BenchmarkSetContains_Generic(b *testing.B) {
	set := NewSet[string]()
	items := generateStringSet(1000)
	for _, item := range items {
		set.Add(item)
	}
	testItem := "z"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(testItem)
	}
}

// Benchmark Set Contains - 专用 string 版本
func BenchmarkSetContains_String(b *testing.B) {
	set := NewStringSet()
	items := generateStringSet(1000)
	for _, item := range items {
		set.Add(item)
	}
	testItem := "z"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(testItem)
	}
}

// Benchmark Sort
func BenchmarkSort(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = 1000 - i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testSlice := make([]int, len(slice))
		copy(testSlice, slice)
		Sort(testSlice)
	}
}

// Benchmark BinarySearch
func BenchmarkBinarySearch(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(slice, 500)
	}
}

// Benchmark LinearSearch
func BenchmarkLinearSearch(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(slice, 500)
	}
}
