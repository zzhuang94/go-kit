package algo

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	if !set.IsEmpty() {
		t.Error("Set should be empty")
	}
	set.Add(1)
	set.Add(2)
	set.Add(2)
	if set.Size() != 2 {
		t.Errorf("Expected size 2, got %d", set.Size())
	}
	if !set.Contains(1) {
		t.Error("Contains failed")
	}
	set.Remove(1)
	if set.Contains(1) {
		t.Error("Remove failed")
	}
	slice := set.ToSlice()
	if len(slice) != 1 {
		t.Error("ToSlice failed")
	}
	set.Clear()
	if !set.IsEmpty() {
		t.Error("Clear failed")
	}
}

func TestSetOperations(t *testing.T) {
	set1 := NewSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set2 := NewSet[int]()
	set2.Add(2)
	set2.Add(3)
	set2.Add(4)
	union := set1.Union(set2)
	if union.Size() != 4 {
		t.Errorf("Expected union size 4, got %d", union.Size())
	}
	intersect := set1.Intersect(set2)
	if intersect.Size() != 2 {
		t.Errorf("Expected intersect size 2, got %d", intersect.Size())
	}
	diff := set1.Diff(set2)
	if diff.Size() != 1 || !diff.Contains(1) {
		t.Error("Diff failed")
	}
}
