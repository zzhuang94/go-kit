package lib

import (
	"testing"
)

func TestChoice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	value, ok := Choice(slice)
	if !ok {
		t.Error("Choice failed")
	}
	found := false
	for _, v := range slice {
		if v == value {
			found = true
			break
		}
	}
	if !found {
		t.Error("Choice returned invalid value")
	}
	_, ok = Choice([]int{})
	if ok {
		t.Error("Choice should return false for empty slice")
	}
}

func TestShuffle(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := Shuffle(slice)
	if len(result) != len(slice) {
		t.Error("Shuffle length failed")
	}
	same := true
	for i := range slice {
		if slice[i] != result[i] {
			same = false
			break
		}
	}
	if same {
		t.Error("Shuffle may not have shuffled")
	}
}
