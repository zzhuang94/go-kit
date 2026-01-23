# Algorithms Package

Provides common data structures and algorithm implementations with generics support.

## 📊 Data Structures

### Stack

- `NewStack[T any]() *Stack[T]` - Create new stack
- `Push(item T)` - Push item onto stack
- `Pop() (T, bool)` - Pop item from stack
- `Peek() (T, bool)` - Peek at top element
- `Size() int` - Return stack size
- `IsEmpty() bool` - Check if stack is empty
- `Clear()` - Clear stack

### Queue

- `NewQueue[T any]() *Queue[T]` - Create new queue
- `Enqueue(item T)` - Enqueue item
- `Dequeue() (T, bool)` - Dequeue item
- `Peek() (T, bool)` - Peek at front element
- `Size() int` - Return queue size
- `IsEmpty() bool` - Check if queue is empty
- `Clear()` - Clear queue

### Set

- `NewSet[T comparable]() *Set[T]` - Create new set
- `Add(item T)` - Add element
- `Remove(item T)` - Remove element
- `Contains(item T) bool` - Check if contains element
- `Size() int` - Return set size
- `IsEmpty() bool` - Check if set is empty
- `Clear()` - Clear set
- `ToSlice() []T` - Convert to slice
- `Union(other *Set[T]) *Set[T]` - Get union
- `Intersect(other *Set[T]) *Set[T]` - Get intersection
- `Diff(other *Set[T]) *Set[T]` - Get difference

### LRU Cache

- `NewLRU[K comparable, V any](capacity int) *LRU[K, V]` - Create LRU cache
- `Get(key K) (V, bool)` - Get value
- `Set(key K, value V)` - Set value
- `Remove(key K) bool` - Remove key-value pair
- `Clear()` - Clear cache
- `Size() int` - Return current size
- `Cap() int` - Return capacity

## 🔢 Algorithms

### Sorting Algorithms

- `Sort[T cmp.Ordered](slice []T)` - Sort slice using quicksort
- `SortDesc[T cmp.Ordered](slice []T)` - Sort slice in descending order
- `SortFunc[T any](slice []T, less func(a, b T) bool)` - Sort slice using custom comparison function
- `IsSorted[T cmp.Ordered](slice []T) bool` - Check if slice is sorted
- `IsSortedFunc[T any](slice []T, less func(a, b T) bool) bool` - Check if slice is sorted using custom comparison function

### Search Algorithms

- `BinarySearch[T cmp.Ordered](slice []T, target T) int` - Binary search, return index of target, -1 if not found
- `BinarySearchFunc[T any](slice []T, target T, cmp func(a, b T) int) int` - Binary search with custom comparison function
- `LinearSearch[T comparable](slice []T, target T) int` - Linear search, return index of target, -1 if not found
- `FindFirst[T any](slice []T, fn func(T) bool) int` - Find first element index matching condition
- `FindLast[T any](slice []T, fn func(T) bool) int` - Find last element index matching condition
- `LowerBound[T cmp.Ordered](slice []T, target T) int` - Find first element index >= target
- `UpperBound[T cmp.Ordered](slice []T, target T) int` - Find first element index > target

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/algo"
)

func main() {
	// Stack
	stack := algo.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	value, _ := stack.Pop()
	fmt.Println(value) // 2
	
	// Queue
	queue := algo.NewQueue[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	value, _ := queue.Dequeue()
	fmt.Println(value) // "first"
	
	// Set
	set := algo.NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(2)
	fmt.Println(set.Size()) // 2
	fmt.Println(set.Contains(1)) // true
	
	// LRU Cache
	lru := algo.NewLRU[string, int](3)
	lru.Set("a", 1)
	lru.Set("b", 2)
	value, _ = lru.Get("a")
	fmt.Println(value) // 1
	
	// Sorting
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	algo.Sort(numbers)
	fmt.Println(numbers) // [1, 1, 2, 3, 4, 5, 6, 9]
	
	// Binary Search
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := algo.BinarySearch(sorted, 5)
	fmt.Println(index) // 4
	
	// Linear Search
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6}
	index = algo.LinearSearch(slice, 5)
	fmt.Println(index) // 4
}
```

## Running Tests

```bash
go test ./algo -v
```
