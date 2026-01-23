# Rand 工具包

提供常用的随机数生成工具函数。

## 随机数生成

- `Int() int` - 生成随机整数
- `IntRange(min, max int) int` - 生成指定范围的随机整数 [min, max)
- `Float64() float64` - 生成随机浮点数 [0.0, 1.0)

## 随机字符串

- `String(length int) string` - 生成随机字符串（字母数字）
- `StringWithCharset(length int, charset string) string` - 使用指定字符集生成随机字符串

## 随机字节数组

- `Bytes(length int) []byte` - 生成随机字节数组

## 随机选择

- `Choice[T any](slice []T) (T, bool)` - 从切片中随机选择一个元素
- `Shuffle[T any](slice []T) []T` - 随机打乱切片

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/rand"
)

func main() {
	// 生成随机整数
	fmt.Println(rand.IntRange(1, 100))
	
	// 生成随机字符串
	fmt.Println(rand.String(10))
	
	// 生成指定字符集的随机字符串
	fmt.Println(rand.StringWithCharset(10, "0123456789"))
	
	// 从切片中随机选择
	slice := []string{"a", "b", "c"}
	value, _ := rand.Choice(slice)
	fmt.Println(value)
	
	// 随机打乱切片
	shuffled := rand.Shuffle(slice)
	fmt.Println(shuffled)
}
```

## 运行测试

```bash
go test ./rand -v
```
