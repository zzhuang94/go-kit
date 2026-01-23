# Convert 工具包

提供常用的类型转换工具函数。

## 字符串转其他类型

- `ToInt(s string) (int, error)` - 字符串转整数
- `ToInt64(s string) (int64, error)` - 字符串转int64
- `ToFloat64(s string) (float64, error)` - 字符串转float64
- `ToBool(s string) (bool, error)` - 字符串转布尔值

## 其他类型转字符串

- `IntToString(i int) string` - 整数转字符串
- `Int64ToString(i int64) string` - int64转字符串
- `Float64ToString(f float64) string` - float64转字符串
- `BoolToString(b bool) string` - 布尔值转字符串

## 字节数组和字符串互转

- `BytesToString(b []byte) string` - 字节数组转字符串
- `StringToBytes(s string) []byte` - 字符串转字节数组

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/convert"
)

func main() {
	// 字符串转数字
	num, _ := convert.ToInt("123")
	fmt.Println(num)
	
	// 数字转字符串
	str := convert.IntToString(123)
	fmt.Println(str)
	
	// 字符串转布尔值
	b, _ := convert.ToBool("true")
	fmt.Println(b)
	
	// 字节数组和字符串互转
	data := convert.StringToBytes("hello")
	fmt.Println(data)
	str2 := convert.BytesToString(data)
	fmt.Println(str2)
}
```

## 运行测试

```bash
go test ./convert -v
```
