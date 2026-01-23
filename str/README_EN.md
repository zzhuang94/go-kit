# String Utilities Package

Provides common string manipulation utility functions.

## Basic Operations

- `Trim(s string) string` - Trim whitespace from both ends of string
- `Substring(s string, start, end int) string` - Extract substring (supports negative indices)
- `ReplaceAll(s, old, new string) string` - Replace all occurrences of string
- `ReplaceN(s, old, new string, n int) string` - Replace first N occurrences of string
- `Contains(s, substr string) bool` - Check if string contains substring
- `ContainsIgnoreCase(s, substr string) bool` - Check if string contains substring (case insensitive)
- `StartsWith(s, prefix string) bool` - Check if string starts with prefix
- `EndsWith(s, suffix string) bool` - Check if string ends with suffix
- `Repeat(s string, count int) string` - Repeat string specified number of times
- `Reverse(s string) string` - Reverse string
- `PadLeft(s string, length int, pad string) string` - Pad string to specified length on the left
- `PadRight(s string, length int, pad string) string` - Pad string to specified length on the right
- `Truncate(s string, length int) string` - Truncate string to specified length
- `TruncateWithEllipsis(s string, length int) string` - Truncate string and add ellipsis

## Formatting

- `CamelCase(s string) string` - Convert to camelCase
- `SnakeCase(s string) string` - Convert to snake_case
- `KebabCase(s string) string` - Convert to kebab-case
- `TitleCase(s string) string` - Convert to Title Case (first letter of each word capitalized)
- `Pluralize(s string) string` - Simple pluralization (basic rules)

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
	fmt.Println(str.Substring("Hello, World", 0, 5)) // "Hello"
	fmt.Println(str.Reverse("hello"))                // "olleh"
	
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
