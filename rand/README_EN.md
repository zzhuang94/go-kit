# Random Number Utilities Package

Provides common random number generation utility functions.

## Random Number Generation

- `Int() int` - Generate random integer
- `IntRange(min, max int) int` - Generate random integer in range [min, max)
- `Float64() float64` - Generate random float [0.0, 1.0)

## Random String

- `String(length int) string` - Generate random string (alphanumeric)
- `StringWithCharset(length int, charset string) string` - Generate random string with specified charset

## Random Byte Array

- `Bytes(length int) []byte` - Generate random byte array

## Random Selection

- `Choice[T any](slice []T) (T, bool)` - Randomly select an element from slice
- `Shuffle[T any](slice []T) []T` - Randomly shuffle slice

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/rand"
)

func main() {
	// Generate random integer
	fmt.Println(rand.IntRange(1, 100))
	
	// Generate random string
	fmt.Println(rand.String(10))
	
	// Generate random string with specified charset
	fmt.Println(rand.StringWithCharset(10, "0123456789"))
	
	// Randomly select from slice
	slice := []string{"a", "b", "c"}
	value, _ := rand.Choice(slice)
	fmt.Println(value)
	
	// Randomly shuffle slice
	shuffled := rand.Shuffle(slice)
	fmt.Println(shuffled)
}
```

## Running Tests

```bash
go test ./rand -v
```
