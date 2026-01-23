# Data Structures Package

Provides common data structure implementations with generics support.

## Stack

- `NewStack[T any]() *Stack[T]` - Create new stack
- `Push(item T)` - Push item onto stack
- `Pop() (T, bool)` - Pop item from stack
- `Peek() (T, bool)` - Peek at top element
- `Size() int` - Return stack size
- `IsEmpty() bool` - Check if stack is empty
- `Clear()` - Clear stack

## Queue

- `NewQueue[T any]() *Queue[T]` - Create new queue
- `Enqueue(item T)` - Enqueue item
- `Dequeue() (T, bool)` - Dequeue item
- `Peek() (T, bool)` - Peek at front element
- `Size() int` - Return queue size
- `IsEmpty() bool` - Check if queue is empty
- `Clear()` - Clear queue

## Set

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

## LRU Cache

- `NewLRU[K comparable, V any](capacity int) *LRU[K, V]` - Create LRU cache
- `Get(key K) (V, bool)` - Get value
- `Set(key K, value V)` - Set value
- `Remove(key K) bool` - Remove key-value pair
- `Clear()` - Clear cache
- `Size() int` - Return current size
- `Cap() int` - Return capacity

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/structs"
)

func main() {
	// Stack
	stack := structs.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	value, _ := stack.Pop()
	fmt.Println(value) // 2
	
	// Queue
	queue := structs.NewQueue[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	value, _ := queue.Dequeue()
	fmt.Println(value) // "first"
	
	// Set
	set := structs.NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(2)
	fmt.Println(set.Size()) // 2
	fmt.Println(set.Contains(1)) // true
	
	// LRU Cache
	lru := structs.NewLRU[string, int](3)
	lru.Set("a", 1)
	lru.Set("b", 2)
	value, _ = lru.Get("a")
	fmt.Println(value) // 1
}
```

## Running Tests

```bash
go test ./structs -v
```
