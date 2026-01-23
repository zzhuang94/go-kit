package convert

import (
	"testing"
)

func TestToInt(t *testing.T) {
	result, err := ToInt("123")
	if err != nil {
		t.Fatalf("ToInt failed: %v", err)
	}
	if result != 123 {
		t.Errorf("Expected 123, got %d", result)
	}
}

func TestToInt64(t *testing.T) {
	result, err := ToInt64("123456789012345")
	if err != nil {
		t.Fatalf("ToInt64 failed: %v", err)
	}
	if result != 123456789012345 {
		t.Error("ToInt64 failed")
	}
}

func TestToFloat64(t *testing.T) {
	result, err := ToFloat64("123.456")
	if err != nil {
		t.Fatalf("ToFloat64 failed: %v", err)
	}
	if result != 123.456 {
		t.Errorf("Expected 123.456, got %f", result)
	}
}

func TestToBool(t *testing.T) {
	result, err := ToBool("true")
	if err != nil {
		t.Fatalf("ToBool failed: %v", err)
	}
	if !result {
		t.Error("ToBool failed")
	}
}

func TestIntToString(t *testing.T) {
	result := IntToString(123)
	if result != "123" {
		t.Errorf("Expected '123', got %q", result)
	}
}

func TestInt64ToString(t *testing.T) {
	result := Int64ToString(123456789012345)
	if result != "123456789012345" {
		t.Error("Int64ToString failed")
	}
}

func TestFloat64ToString(t *testing.T) {
	result := Float64ToString(123.456)
	if result != "123.456" {
		t.Errorf("Expected '123.456', got %q", result)
	}
}

func TestBoolToString(t *testing.T) {
	if BoolToString(true) != "true" {
		t.Error("BoolToString failed")
	}
	if BoolToString(false) != "false" {
		t.Error("BoolToString failed")
	}
}

func TestBytesToString(t *testing.T) {
	data := []byte("hello")
	result := BytesToString(data)
	if result != "hello" {
		t.Errorf("Expected 'hello', got %q", result)
	}
}

func TestStringToBytes(t *testing.T) {
	s := "hello"
	result := StringToBytes(s)
	if string(result) != s {
		t.Errorf("Expected %q, got %q", s, string(result))
	}
}
