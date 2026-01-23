# JSON Utilities Package

Provides common JSON operation utility functions.

## Serialization and Deserialization

- `Marshal(v interface{}) ([]byte, error)` - JSON serialization
- `Unmarshal(data []byte, v interface{}) error` - JSON deserialization
- `MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)` - JSON serialization (formatted)
- `PrettyPrint(v interface{}) error` - Pretty print JSON

## JSON Operations

- `Get(data []byte, key string) (interface{}, error)` - Get JSON field value
- `Set(data []byte, key string, value interface{}) ([]byte, error)` - Set JSON field value
- `Merge(data1, data2 []byte) ([]byte, error)` - Merge JSON objects

## File Operations

- `ReadFile(path string, v interface{}) error` - Read JSON from file
- `WriteFile(path string, v interface{}) error` - Write JSON to file

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/json"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	
	// Serialization
	p := Person{Name: "Alice", Age: 30}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	
	// Deserialization
	var p2 Person
	json.Unmarshal(data, &p2)
	fmt.Println(p2)
	
	// Pretty print
	json.PrettyPrint(p)
	
	// Get field value
	value, _ := json.Get(data, "name")
	fmt.Println(value)
	
	// Set field value
	newData, _ := json.Set(data, "age", 31)
	fmt.Println(string(newData))
}
```

## Running Tests

```bash
go test ./json -v
```
