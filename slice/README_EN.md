# Slice Utilities Package

Provides common slice operation utility functions with generics support.

## Basic Operations

- `Contains[T comparable](slice []T, item T) bool` - Check if slice contains element
- `IndexOf[T comparable](slice []T, item T) int` - Find element index in slice
- `LastIndexOf[T comparable](slice []T, item T) int` - Find element index from end
- `Unique[T comparable](slice []T) []T` - Remove duplicate elements
- `Reverse[T any](slice []T) []T` - Reverse slice
- `Shuffle[T any](slice []T) []T` - Randomly shuffle slice
- `Chunk[T any](slice []T, size int) [][]T` - Split slice into chunks
- `Flatten[T any](slices [][]T) []T` - Flatten nested slices
- `Intersect[T comparable](slice1, slice2 []T) []T` - Get intersection of two slices
- `Union[T comparable](slice1, slice2 []T) []T` - Get union of two slices
- `Diff[T comparable](slice1, slice2 []T) []T` - Get difference of two slices
- `Remove[T comparable](slice []T, item T) []T` - Remove all matching elements
- `RemoveAt[T any](slice []T, index int) []T` - Remove element at index
- `Insert[T any](slice []T, index int, item T) []T` - Insert element at index
- `First[T any](slice []T) T` - Get first element
- `Last[T any](slice []T) T` - Get last element
- `Take[T any](slice []T, n int) []T` - Take first N elements
- `Drop[T any](slice []T, n int) []T` - Drop first N elements

## Filter and Find

- `Filter[T any](slice []T, fn func(T) bool) []T` - Filter slice, keep elements matching condition
- `Find[T any](slice []T, fn func(T) bool) (T, bool)` - Find first element matching condition
- `FindIndex[T any](slice []T, fn func(T) bool) int` - Find first element index matching condition
- `FindLast[T any](slice []T, fn func(T) bool) (T, bool)` - Find last element matching condition
- `FindLastIndex[T any](slice []T, fn func(T) bool) int` - Find last element index matching condition
- `Every[T any](slice []T, fn func(T) bool) bool` - Check if all elements match condition
- `Some[T any](slice []T, fn func(T) bool) bool` - Check if at least one element matches condition
- `Count[T any](slice []T, fn func(T) bool) int` - Count elements matching condition
- `GroupBy[T any, K comparable](slice []T, fn func(T) K) map[K][]T` - Group by function

## Transform

- `Map[T, R any](slice []T, fn func(T) R) []R` - Map transform slice
- `Reduce[T, R any](slice []T, initial R, fn func(R, T) R) R` - Reduce slice
- `FlatMap[T, R any](slice []T, fn func(T) []R) []R` - Flat map
- `Partition[T any](slice []T, fn func(T) bool) ([]T, []T)` - Partition slice into two parts

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/slice"
)

func main() {
	// Basic operations
	numbers := []int{1, 2, 3, 2, 4, 5}
	fmt.Println(slice.Contains(numbers, 3))        // true
	fmt.Println(slice.Unique(numbers))               // [1, 2, 3, 4, 5]
	fmt.Println(slice.Reverse(numbers))              // [5, 4, 2, 3, 2, 1]
	
	// Filter and find
	even := slice.Filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(even)                                // [2, 2, 4]
	
	value, found := slice.Find(numbers, func(x int) bool {
		return x > 3
	})
	fmt.Println(value, found)                       // 4 true
	
	// Transform
	doubled := slice.Map(numbers, func(x int) int {
		return x * 2
	})
	fmt.Println(doubled)                            // [2, 4, 6, 4, 8, 10]
	
	sum := slice.Reduce(numbers, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Println(sum)                                // 17
}
```

## Running Tests

```bash
go test ./slice -v
```
