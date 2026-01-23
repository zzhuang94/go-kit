# Algo 工具包

提供常用的数据结构和算法实现，支持泛型。

## 📊 数据结构

### Stack 栈

- `NewStack[T any]() *Stack[T]` - 创建新栈
- `Push(item T)` - 入栈
- `Pop() (T, bool)` - 出栈
- `Peek() (T, bool)` - 查看栈顶元素
- `Size() int` - 返回栈大小
- `IsEmpty() bool` - 检查栈是否为空
- `Clear()` - 清空栈

### Queue 队列

- `NewQueue[T any]() *Queue[T]` - 创建新队列
- `Enqueue(item T)` - 入队
- `Dequeue() (T, bool)` - 出队
- `Peek() (T, bool)` - 查看队首元素
- `Size() int` - 返回队列大小
- `IsEmpty() bool` - 检查队列是否为空
- `Clear()` - 清空队列

### Set 集合

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

### LRU Cache

- `NewLRU[K comparable, V any](capacity int) *LRU[K, V]` - 创建LRU缓存
- `Get(key K) (V, bool)` - 获取值
- `Set(key K, value V)` - 设置值
- `Remove(key K) bool` - 删除键值对
- `Clear()` - 清空缓存
- `Size() int` - 返回当前大小
- `Cap() int` - 返回容量

## 🔢 算法

### 排序算法

- `Sort[T cmp.Ordered](slice []T)` - 对切片进行排序（快速排序）
- `SortDesc[T cmp.Ordered](slice []T)` - 对切片进行降序排序
- `SortFunc[T any](slice []T, less func(a, b T) bool)` - 使用自定义比较函数排序
- `IsSorted[T cmp.Ordered](slice []T) bool` - 检查切片是否已排序
- `IsSortedFunc[T any](slice []T, less func(a, b T) bool) bool` - 使用自定义比较函数检查是否已排序

### 搜索算法

- `BinarySearch[T cmp.Ordered](slice []T, target T) int` - 二分查找，返回索引，不存在返回 -1
- `BinarySearchFunc[T any](slice []T, target T, cmp func(a, b T) int) int` - 使用自定义比较函数进行二分查找
- `LinearSearch[T comparable](slice []T, target T) int` - 线性查找，返回索引，不存在返回 -1
- `FindFirst[T any](slice []T, fn func(T) bool) int` - 查找第一个满足条件的元素索引
- `FindLast[T any](slice []T, fn func(T) bool) int` - 查找最后一个满足条件的元素索引
- `LowerBound[T cmp.Ordered](slice []T, target T) int` - 查找第一个大于等于目标值的元素索引
- `UpperBound[T cmp.Ordered](slice []T, target T) int` - 查找第一个大于目标值的元素索引

## 使用示例

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
	
	// 排序
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	algo.Sort(numbers)
	fmt.Println(numbers) // [1, 1, 2, 3, 4, 5, 6, 9]
	
	// 二分查找
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := algo.BinarySearch(sorted, 5)
	fmt.Println(index) // 4
	
	// 线性查找
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6}
	index = algo.LinearSearch(slice, 5)
	fmt.Println(index) // 4
}
```

## 运行测试

```bash
go test ./algo -v
```
