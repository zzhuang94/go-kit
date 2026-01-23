package algo

import (
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[string, int](3)
	if lru.Cap() != 3 {
		t.Error("Cap failed")
	}
	lru.Set("a", 1)
	lru.Set("b", 2)
	lru.Set("c", 3)
	if lru.Size() != 3 {
		t.Error("Size failed")
	}
	value, ok := lru.Get("a")
	if !ok || value != 1 {
		t.Error("Get failed")
	}
	lru.Set("d", 4)
	if lru.Size() != 3 {
		t.Error("Size after capacity exceeded failed")
	}
	_, ok = lru.Get("b")
	if ok {
		t.Error("LRU eviction failed")
	}
	lru.Remove("a")
	if lru.Size() != 2 {
		t.Error("Remove failed")
	}
	lru.Clear()
	if lru.Size() != 0 {
		t.Error("Clear failed")
	}
}
