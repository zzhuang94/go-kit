# Structs 工具包

提供常用的数据结构实现，支持泛型。

## Stack 栈

- `NewStack[T any]() *Stack[T]` - 创建新栈
- `Push(item T)` - 入栈
- `Pop() (T, bool)` - 出栈
- `Peek() (T, bool)` - 查看栈顶元素
- `Size() int` - 返回栈大小
- `IsEmpty() bool` - 检查栈是否为空
- `Clear()` - 清空栈

## Queue 队列

- `NewQueue[T any]() *Queue[T]` - 创建新队列
- `Enqueue(item T)` - 入队
- `Dequeue() (T, bool)` - 出队
- `Peek() (T, bool)` - 查看队首元素
- `Size() int` - 返回队列大小
- `IsEmpty() bool` - 检查队列是否为空
- `Clear()` - 清空队列

## Set 集合

- `NewSet[T comparable]() *Set[T]` - 创建新集合
- `Add(item T)` - 添加元素
- `Remove(item T)` - 删除元素
- `Contains(item T) bool` - 检查是否包含元素
- `Size() int` - 返回集合大小
- `IsEmpty() bool` - 检查集合是否为空
- `Clear()` - 清空集合
- `ToSlice() []T` - 转换为切片
- `Union(other *Set[T]) *Set[T]` - 求并集
- `Intersect(other *Set[T]) *Set[T]` - 求交集
- `Diff(other *Set[T]) *Set[T]` - 求差集

## LRU Cache

- `NewLRU[K comparable, V any](capacity int) *LRU[K, V]` - 创建LRU缓存
- `Get(key K) (V, bool)` - 获取值
- `Set(key K, value V)` - 设置值
- `Remove(key K) bool` - 删除键值对
- `Clear()` - 清空缓存
- `Size() int` - 返回当前大小
- `Cap() int` - 返回容量

## 使用示例

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

## 运行测试

```bash
go test ./structs -v
```
