# Time Utilities Package

Provides common time operation utility functions.

## Formatting

- `Format(t time.Time, layout string) string` - Format time to string
- `FormatDate(t time.Time) string` - Format date (YYYY-MM-DD)
- `FormatTime(t time.Time) string` - Format time (HH:mm:ss)
- `FormatDateTime(t time.Time) string` - Format date and time (YYYY-MM-DD HH:mm:ss)
- `FormatTimestamp(t time.Time) int64` - Format timestamp

## Parsing

- `Parse(layout, value string) (time.Time, error)` - Parse time string
- `ParseDate(value string) (time.Time, error)` - Parse date string (YYYY-MM-DD)
- `ParseTime(value string) (time.Time, error)` - Parse time string (HH:mm:ss)
- `ParseDateTime(value string) (time.Time, error)` - Parse date and time string (YYYY-MM-DD HH:mm:ss)

## Calculations

- `AddDays(t time.Time, days int) time.Time` - Add days
- `AddHours(t time.Time, hours int) time.Time` - Add hours
- `AddMinutes(t time.Time, minutes int) time.Time` - Add minutes
- `DiffDays(t1, t2 time.Time) int` - Calculate days difference
- `DiffHours(t1, t2 time.Time) int64` - Calculate hours difference
- `DiffMinutes(t1, t2 time.Time) int64` - Calculate minutes difference
- `IsToday(t time.Time) bool` - Check if time is today
- `IsYesterday(t time.Time) bool` - Check if time is yesterday
- `IsTomorrow(t time.Time) bool` - Check if time is tomorrow
- `StartOfDay(t time.Time) time.Time` - Get start of day
- `EndOfDay(t time.Time) time.Time` - Get end of day
- `StartOfWeek(t time.Time) time.Time` - Get start of week (Monday)
- `StartOfMonth(t time.Time) time.Time` - Get start of month
- `StartOfYear(t time.Time) time.Time` - Get start of year

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/time"
	"time"
)

func main() {
	now := time.Now()
	
	// Formatting
	fmt.Println(time.FormatDate(now))        // "2024-01-15"
	fmt.Println(time.FormatDateTime(now))    // "2024-01-15 10:30:45"
	
	// Parsing
	tm, _ := time.ParseDate("2024-01-15")
	fmt.Println(tm)
	
	// Calculations
	tomorrow := time.AddDays(now, 1)
	diff := time.DiffDays(tomorrow, now)
	fmt.Println(diff) // 1
	
	// Check
	fmt.Println(time.IsToday(now))           // true
	fmt.Println(time.IsTomorrow(tomorrow))   // true
	
	// Get start time
	start := time.StartOfDay(now)
	fmt.Println(start)
}
```

## Running Tests

```bash
go test ./time -v
```
