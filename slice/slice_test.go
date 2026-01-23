package slice

import (
	"testing"
)

func TestContains(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	if !Contains(slice, 3) {
		t.Error("Contains failed")
	}
	if Contains(slice, 6) {
		t.Error("Contains failed")
	}
}

func TestIndexOf(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if IndexOf(slice, "b") != 1 {
		t.Error("IndexOf failed")
	}
	if IndexOf(slice, "d") != -1 {
		t.Error("IndexOf failed")
	}
}

func TestLastIndexOf(t *testing.T) {
	slice := []int{1, 2, 3, 2, 4}
	if LastIndexOf(slice, 2) != 3 {
		t.Error("LastIndexOf failed")
	}
}

func TestUnique(t *testing.T) {
	slice := []int{1, 2, 2, 3, 3, 3}
	result := Unique(slice)
	if len(result) != 3 {
		t.Errorf("Expected 3 unique elements, got %d", len(result))
	}
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	result := Reverse(slice)
	if result[0] != 4 || result[3] != 1 {
		t.Error("Reverse failed")
	}
}

func TestChunk(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	chunks := Chunk(slice, 2)
	if len(chunks) != 3 {
		t.Errorf("Expected 3 chunks, got %d", len(chunks))
	}
	if len(chunks[0]) != 2 {
		t.Error("Chunk size failed")
	}
}

func TestFlatten(t *testing.T) {
	slices := [][]int{{1, 2}, {3, 4}, {5}}
	result := Flatten(slices)
	if len(result) != 5 {
		t.Errorf("Expected 5 elements, got %d", len(result))
	}
}

func TestIntersect(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
	result := Intersect(slice1, slice2)
	if len(result) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result))
	}
}

func TestUnion(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5}
	result := Union(slice1, slice2)
	if len(result) != 5 {
		t.Errorf("Expected 5 elements, got %d", len(result))
	}
}

func TestDiff(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4}
	result := Diff(slice1, slice2)
	if len(result) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result))
	}
}

func TestRemove(t *testing.T) {
	slice := []int{1, 2, 3, 2, 4}
	result := Remove(slice, 2)
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}
}

func TestRemoveAt(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	result := RemoveAt(slice, 1)
	if len(result) != 3 || result[1] != 3 {
		t.Error("RemoveAt failed")
	}
}

func TestInsert(t *testing.T) {
	slice := []int{1, 2, 4}
	result := Insert(slice, 2, 3)
	if len(result) != 4 || result[2] != 3 {
		t.Error("Insert failed")
	}
}

func TestFirst(t *testing.T) {
	slice := []int{1, 2, 3}
	if First(slice) != 1 {
		t.Error("First failed")
	}
}

func TestLast(t *testing.T) {
	slice := []int{1, 2, 3}
	if Last(slice) != 3 {
		t.Error("Last failed")
	}
}

func TestTake(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := Take(slice, 3)
	if len(result) != 3 {
		t.Error("Take failed")
	}
}

func TestDrop(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := Drop(slice, 2)
	if len(result) != 3 || result[0] != 3 {
		t.Error("Drop failed")
	}
}

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := Filter(slice, func(x int) bool {
		return x%2 == 0
	})
	if len(result) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result))
	}
}

func TestFind(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	value, found := Find(slice, func(x int) bool {
		return x > 3
	})
	if !found || value != 4 {
		t.Error("Find failed")
	}
}

func TestFindIndex(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	index := FindIndex(slice, func(x int) bool {
		return x > 3
	})
	if index != 3 {
		t.Errorf("Expected index 3, got %d", index)
	}
}

func TestEvery(t *testing.T) {
	slice := []int{2, 4, 6}
	if !Every(slice, func(x int) bool {
		return x%2 == 0
	}) {
		t.Error("Every failed")
	}
}

func TestSome(t *testing.T) {
	slice := []int{1, 2, 3}
	if !Some(slice, func(x int) bool {
		return x%2 == 0
	}) {
		t.Error("Some failed")
	}
}

func TestCount(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	count := Count(slice, func(x int) bool {
		return x%2 == 0
	})
	if count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
}

func TestGroupBy(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	groups := GroupBy(slice, func(x int) string {
		if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if len(groups["even"]) != 2 {
		t.Errorf("Expected 2 even numbers, got %d", len(groups["even"]))
	}
}

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	result := Map(slice, func(x int) int {
		return x * 2
	})
	if result[0] != 2 || result[1] != 4 {
		t.Error("Map failed")
	}
}

func TestReduce(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	sum := Reduce(slice, 0, func(acc, x int) int {
		return acc + x
	})
	if sum != 10 {
		t.Errorf("Expected sum 10, got %d", sum)
	}
}

func TestFlatMap(t *testing.T) {
	slice := []int{1, 2, 3}
	result := FlatMap(slice, func(x int) []int {
		return []int{x, x * 2}
	})
	if len(result) != 6 {
		t.Errorf("Expected 6 elements, got %d", len(result))
	}
}

func TestPartition(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	even, odd := Partition(slice, func(x int) bool {
		return x%2 == 0
	})
	if len(even) != 2 || len(odd) != 3 {
		t.Errorf("Partition failed: even=%d, odd=%d", len(even), len(odd))
	}
}
