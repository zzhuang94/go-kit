package algo

import (
	"testing"
)

func TestSort(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6}
	Sort(slice)
	if !IsSorted(slice) {
		t.Error("Sort failed")
	}
	expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
	for i, v := range slice {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestSortDesc(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6}
	SortDesc(slice)
	expected := []int{9, 6, 5, 4, 3, 2, 1, 1}
	for i, v := range slice {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestSortFunc(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6}
	SortFunc(slice, func(a, b int) bool {
		return a < b
	})
	if !IsSorted(slice) {
		t.Error("SortFunc failed")
	}
}

func TestIsSorted(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	if !IsSorted(sorted) {
		t.Error("IsSorted failed for sorted slice")
	}
	unsorted := []int{3, 1, 4, 2, 5}
	if IsSorted(unsorted) {
		t.Error("IsSorted failed for unsorted slice")
	}
}

func TestIsSortedFunc(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	if !IsSortedFunc(sorted, func(a, b int) bool {
		return a < b
	}) {
		t.Error("IsSortedFunc failed for sorted slice")
	}
	unsorted := []int{3, 1, 4, 2, 5}
	if IsSortedFunc(unsorted, func(a, b int) bool {
		return a < b
	}) {
		t.Error("IsSortedFunc failed for unsorted slice")
	}
}
