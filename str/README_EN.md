# String Utilities Package

Provides common string manipulation utility functions.

## Basic Operations

- `Trim(s string) string` - Trim whitespace from both ends of string
- `Sub(s string, start, end int) string` - Extract substring (supports negative indices)
- `Reverse(s string) string` - Reverse string
- `PadLeft(s string, length int, pad string) string` - Pad string to specified length on the left
- `PadRight(s string, length int, pad string) string` - Pad string to specified length on the right
- `Truncate(s string, length int) string` - Truncate string to specified length
- `SplitLines(str string) []string` - Split string by lines, remove empty lines and duplicates, trim whitespace
- `Md5(params any) string` - Calculate MD5 hash of parameters (after JSON serialization)
- `Uuid() string` - Generate UUID v4 string

## Formatting

- `CamelCase(s string) string` - Convert to camelCase
- `SnakeCase(s string) string` - Convert to snake_case
- `KebabCase(s string) string` - Convert to kebab-case
- `TitleCase(s string) string` - Convert to Title Case (first letter of each word capitalized)
- `Pluralize(s string) string` - Simple pluralization (basic rules)
- `ParseAndFormatJson(str string) (string, error)` - Parse and format JSON string

## Validation

- `IsEmail(s string) bool` - Validate email format
- `IsURL(s string) bool` - Validate URL format
- `IsIP(s string) bool` - Validate IP address format (IPv4)
- `IsPhone(s string) bool` - Validate phone number format (Mainland China)
- `IsNumeric(s string) bool` - Validate if string is numeric
- `IsAlpha(s string) bool` - Validate if string is alphabetic
- `IsAlphaNumeric(s string) bool` - Validate if string is alphanumeric
- `IsEmpty(s string) bool` - Check if string is empty
- `IsBlank(s string) bool` - Check if string is blank (empty or only whitespace)

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/str"
)

func main() {
	// Basic operations
	fmt.Println(str.Trim("  hello  "))              // "hello"
	fmt.Println(str.Sub("Hello, World", 0, 5))      // "Hello"
	fmt.Println(str.Reverse("hello"))                // "olleh"
	fmt.Println(str.Md5("hello"))                    // "5d41402abc4b2a76b9719d911017c592"
	fmt.Println(str.Uuid())                          // "a1b2c3d4e5f6..."
	
	// Formatting
	fmt.Println(str.CamelCase("hello_world"))       // "helloWorld"
	fmt.Println(str.SnakeCase("HelloWorld"))         // "hello_world"
	fmt.Println(str.KebabCase("HelloWorld"))         // "hello-world"
	
	// Validation
	fmt.Println(str.IsEmail("test@example.com"))     // true
	fmt.Println(str.IsURL("https://example.com"))    // true
	fmt.Println(str.IsPhone("13800138000"))          // true
}
```

## Running Tests

```bash
go test ./str -v
```
