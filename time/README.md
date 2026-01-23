# Time 工具包

提供常用的时间操作工具函数。

## 格式化

- `Format(t time.Time, layout string) string` - 格式化时间为字符串
- `FormatDate(t time.Time) string` - 格式化日期（YYYY-MM-DD）
- `FormatTime(t time.Time) string` - 格式化时间（HH:mm:ss）
- `FormatDateTime(t time.Time) string` - 格式化日期时间（YYYY-MM-DD HH:mm:ss）
- `FormatTimestamp(t time.Time) int64` - 格式化时间戳

## 解析

- `Parse(layout, value string) (time.Time, error)` - 解析时间字符串
- `ParseDate(value string) (time.Time, error)` - 解析日期字符串（YYYY-MM-DD）
- `ParseTime(value string) (time.Time, error)` - 解析时间字符串（HH:mm:ss）
- `ParseDateTime(value string) (time.Time, error)` - 解析日期时间字符串（YYYY-MM-DD HH:mm:ss）

## 计算

- `AddDays(t time.Time, days int) time.Time` - 添加天数
- `AddHours(t time.Time, hours int) time.Time` - 添加小时
- `AddMinutes(t time.Time, minutes int) time.Time` - 添加分钟
- `DiffDays(t1, t2 time.Time) int` - 计算天数差
- `DiffHours(t1, t2 time.Time) int64` - 计算小时差
- `DiffMinutes(t1, t2 time.Time) int64` - 计算分钟差
- `IsToday(t time.Time) bool` - 判断是否为今天
- `IsYesterday(t time.Time) bool` - 判断是否为昨天
- `IsTomorrow(t time.Time) bool` - 判断是否为明天
- `StartOfDay(t time.Time) time.Time` - 获取一天的开始时间
- `EndOfDay(t time.Time) time.Time` - 获取一天的结束时间
- `StartOfWeek(t time.Time) time.Time` - 获取一周的开始时间（周一）
- `StartOfMonth(t time.Time) time.Time` - 获取一月的开始时间
- `StartOfYear(t time.Time) time.Time` - 获取一年的开始时间

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/time"
	"time"
)

func main() {
	now := time.Now()
	
	// 格式化
	fmt.Println(time.FormatDate(now))        // "2024-01-15"
	fmt.Println(time.FormatDateTime(now))    // "2024-01-15 10:30:45"
	
	// 解析
	tm, _ := time.ParseDate("2024-01-15")
	fmt.Println(tm)
	
	// 计算
	tomorrow := time.AddDays(now, 1)
	diff := time.DiffDays(tomorrow, now)
	fmt.Println(diff) // 1
	
	// 判断
	fmt.Println(time.IsToday(now))           // true
	fmt.Println(time.IsTomorrow(tomorrow))   // true
	
	// 获取开始时间
	start := time.StartOfDay(now)
	fmt.Println(start)
}
```

## 运行测试

```bash
go test ./time -v
```
