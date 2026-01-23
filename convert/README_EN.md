# Convert Utilities Package

Provides common type conversion utility functions.

## String to Other Types

- `ToInt(s string) (int, error)` - Convert string to integer
- `ToInt64(s string) (int64, error)` - Convert string to int64
- `ToFloat64(s string) (float64, error)` - Convert string to float64
- `ToBool(s string) (bool, error)` - Convert string to boolean

## Other Types to String

- `IntToString(i int) string` - Convert integer to string
- `Int64ToString(i int64) string` - Convert int64 to string
- `Float64ToString(f float64) string` - Convert float64 to string
- `BoolToString(b bool) string` - Convert boolean to string

## Byte Array and String Conversion

- `BytesToString(b []byte) string` - Convert byte array to string
- `StringToBytes(s string) []byte` - Convert string to byte array

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/convert"
)

func main() {
	// String to number
	num, _ := convert.ToInt("123")
	fmt.Println(num)
	
	// Number to string
	str := convert.IntToString(123)
	fmt.Println(str)
	
	// String to boolean
	b, _ := convert.ToBool("true")
	fmt.Println(b)
	
	// Byte array and string conversion
	data := convert.StringToBytes("hello")
	fmt.Println(data)
	str2 := convert.BytesToString(data)
	fmt.Println(str2)
}
```

## Running Tests

```bash
go test ./convert -v
```
