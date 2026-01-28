# Slice 工具包

提供常用的切片操作工具函数，支持泛型。

## 基础操作

- `LastIndexOf[T comparable](slice []T, item T) int` - 从后往前查找元素索引
- `Unique[T comparable](slice []T) []T` - 去除重复元素
- `Shuffle[T any](slice []T) []T` - 随机打乱切片
- `Chunk[T any](slice []T, size int) [][]T` - 将切片分块
- `Flatten[T any](slices [][]T) []T` - 扁平化嵌套切片
- `Intersect[T comparable](slice1, slice2 []T) []T` - 求两个切片的交集
- `Union[T comparable](slice1, slice2 []T) []T` - 求两个切片的并集
- `Diff[T comparable](slice1, slice2 []T) []T` - 求两个切片的差集
- `Remove[T comparable](slice []T, item T) []T` - 删除切片中所有匹配的元素
- `RemoveAt[T any](slice []T, index int) []T` - 删除指定索引的元素
- `Insert[T any](slice []T, index int, item T) []T` - 在指定索引插入元素
- `First[T any](slice []T) T` - 获取第一个元素
- `Last[T any](slice []T) T` - 获取最后一个元素
- `Take[T any](slice []T, n int) []T` - 取前N个元素
- `Drop[T any](slice []T, n int) []T` - 跳过前N个元素

## 过滤和查找

- `Filter[T any](slice []T, fn func(T) bool) []T` - 过滤切片，保留满足条件的元素
- `Find[T any](slice []T, fn func(T) bool) (T, bool)` - 查找第一个满足条件的元素
- `FindIndex[T any](slice []T, fn func(T) bool) int` - 查找第一个满足条件的元素索引
- `FindLast[T any](slice []T, fn func(T) bool) (T, bool)` - 从后往前查找第一个满足条件的元素
- `FindLastIndex[T any](slice []T, fn func(T) bool) int` - 从后往前查找第一个满足条件的元素索引
- `Every[T any](slice []T, fn func(T) bool) bool` - 检查是否所有元素都满足条件
- `Some[T any](slice []T, fn func(T) bool) bool` - 检查是否至少有一个元素满足条件
- `Count[T any](slice []T, fn func(T) bool) int` - 统计满足条件的元素个数
- `GroupBy[T any, K comparable](slice []T, fn func(T) K) map[K][]T` - 按指定函数分组

## 转换

- `Map[T, R any](slice []T, fn func(T) R) []R` - 映射转换切片
- `Reduce[T, R any](slice []T, initial R, fn func(R, T) R) R` - 归约切片
- `FlatMap[T, R any](slice []T, fn func(T) []R) []R` - 扁平映射
- `Partition[T any](slice []T, fn func(T) bool) ([]T, []T)` - 将切片分为两部分

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/slice"
)

func main() {
	// 基础操作
	numbers := []int{1, 2, 3, 2, 4, 5}
	fmt.Println(slice.Unique(numbers))               // [1, 2, 3, 4, 5]
	
	// 过滤和查找
	even := slice.Filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(even)                                // [2, 2, 4]
	
	value, found := slice.Find(numbers, func(x int) bool {
		return x > 3
	})
	fmt.Println(value, found)                       // 4 true
	
	// 转换
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

## 运行测试

```bash
go test ./slice -v
```
