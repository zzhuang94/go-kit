package algo

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := BinarySearch(slice, 5)
	if index != 4 {
		t.Errorf("Expected index 4, got %d", index)
	}
	index = BinarySearch(slice, 10)
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestBinarySearchFunc(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := BinarySearchFunc(slice, 5, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	if index != 4 {
		t.Errorf("Expected index 4, got %d", index)
	}
}

func TestLinearSearch(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6}
	index := LinearSearch(slice, 5)
	if index != 4 {
		t.Errorf("Expected index 4, got %d", index)
	}
	index = LinearSearch(slice, 10)
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestFindFirst(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := FindFirst(slice, func(x int) bool {
		return x > 5
	})
	if index != 5 {
		t.Errorf("Expected index 5, got %d", index)
	}
	index = FindFirst(slice, func(x int) bool {
		return x > 10
	})
	if index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestFindLast(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := FindLast(slice, func(x int) bool {
		return x < 5
	})
	if index != 3 {
		t.Errorf("Expected index 3, got %d", index)
	}
}

func TestLowerBound(t *testing.T) {
	slice := []int{1, 2, 3, 3, 3, 4, 5}
	index := LowerBound(slice, 3)
	if index != 2 {
		t.Errorf("Expected index 2, got %d", index)
	}
	index = LowerBound(slice, 0)
	if index != 0 {
		t.Errorf("Expected index 0, got %d", index)
	}
	index = LowerBound(slice, 6)
	if index != len(slice) {
		t.Errorf("Expected index %d, got %d", len(slice), index)
	}
}

func TestUpperBound(t *testing.T) {
	slice := []int{1, 2, 3, 3, 3, 4, 5}
	index := UpperBound(slice, 3)
	if index != 5 {
		t.Errorf("Expected index 5, got %d", index)
	}
	index = UpperBound(slice, 0)
	if index != 0 {
		t.Errorf("Expected index 0, got %d", index)
	}
	index = UpperBound(slice, 6)
	if index != len(slice) {
		t.Errorf("Expected index %d, got %d", len(slice), index)
	}
}
