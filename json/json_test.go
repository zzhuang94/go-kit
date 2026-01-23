package json

import (
	"os"
	"testing"
)

func TestMarshalUnmarshal(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	p := Person{Name: "Alice", Age: 30}
	data, err := Marshal(p)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	var p2 Person
	if err := Unmarshal(data, &p2); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if p2.Name != p.Name || p2.Age != p.Age {
		t.Error("Marshal/Unmarshal failed")
	}
}

func TestMarshalIndent(t *testing.T) {
	m := map[string]interface{}{"name": "Alice", "age": 30}
	data, err := MarshalIndent(m, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent failed: %v", err)
	}
	if len(data) == 0 {
		t.Error("MarshalIndent failed")
	}
}

func TestGet(t *testing.T) {
	data := []byte(`{"name":"Alice","age":30}`)
	value, err := Get(data, "name")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "Alice" {
		t.Errorf("Expected 'Alice', got %v", value)
	}
}

func TestSet(t *testing.T) {
	data := []byte(`{"name":"Alice"}`)
	result, err := Set(data, "age", 30)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}
	value, _ := Get(result, "age")
	if value != float64(30) {
		t.Error("Set failed")
	}
}

func TestMerge(t *testing.T) {
	data1 := []byte(`{"name":"Alice"}`)
	data2 := []byte(`{"age":30}`)
	result, err := Merge(data1, data2)
	if err != nil {
		t.Fatalf("Merge failed: %v", err)
	}
	name, _ := Get(result, "name")
	age, _ := Get(result, "age")
	if name != "Alice" || age != float64(30) {
		t.Error("Merge failed")
	}
}

func TestReadWriteFile(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	p := Person{Name: "Alice", Age: 30}
	tmpFile := "test.json"
	defer os.Remove(tmpFile)
	if err := WriteFile(tmpFile, p); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
	var p2 Person
	if err := ReadFile(tmpFile, &p2); err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	if p2.Name != p.Name || p2.Age != p.Age {
		t.Error("ReadFile/WriteFile failed")
	}
}
