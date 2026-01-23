package rand

import (
	"testing"
)

func TestIntRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := IntRange(1, 10)
		if result < 1 || result >= 10 {
			t.Errorf("Expected range [1, 10), got %d", result)
		}
	}
}

func TestFloat64(t *testing.T) {
	result := Float64()
	if result < 0 || result >= 1 {
		t.Errorf("Expected range [0.0, 1.0), got %f", result)
	}
}

func TestString(t *testing.T) {
	result := String(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
}

func TestStringWithCharset(t *testing.T) {
	charset := "abc"
	result := StringWithCharset(5, charset)
	if len(result) != 5 {
		t.Errorf("Expected length 5, got %d", len(result))
	}
	for _, c := range result {
		if c != 'a' && c != 'b' && c != 'c' {
			t.Errorf("Invalid character: %c", c)
		}
	}
}

func TestBytes(t *testing.T) {
	result := Bytes(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
}

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
