# String 工具包

提供常用的字符串操作工具函数。

## 基础操作

- `Trim(s string) string` - 去除首尾空白字符
- `Substring(s string, start, end int) string` - 截取子串（支持负数索引）
- `ReplaceAll(s, old, new string) string` - 替换所有匹配的字符串
- `ReplaceN(s, old, new string, n int) string` - 替换前N个匹配的字符串
- `Contains(s, substr string) bool` - 检查字符串是否包含子串
- `ContainsIgnoreCase(s, substr string) bool` - 检查字符串是否包含子串（忽略大小写）
- `StartsWith(s, prefix string) bool` - 检查字符串是否以指定前缀开始
- `EndsWith(s, suffix string) bool` - 检查字符串是否以指定后缀结束
- `Repeat(s string, count int) string` - 重复字符串指定次数
- `Reverse(s string) string` - 反转字符串
- `PadLeft(s string, length int, pad string) string` - 左填充字符串到指定长度
- `PadRight(s string, length int, pad string) string` - 右填充字符串到指定长度
- `Truncate(s string, length int) string` - 截断字符串到指定长度
- `TruncateWithEllipsis(s string, length int) string` - 截断字符串并添加省略号

## 格式化

- `CamelCase(s string) string` - 转换为驼峰命名
- `SnakeCase(s string) string` - 转换为蛇形命名
- `KebabCase(s string) string` - 转换为短横线命名
- `TitleCase(s string) string` - 转换为标题格式（每个单词首字母大写）
- `Pluralize(s string) string` - 简单的复数化（基础规则）

## 验证

- `IsEmail(s string) bool` - 验证邮箱格式
- `IsURL(s string) bool` - 验证URL格式
- `IsIP(s string) bool` - 验证IP地址格式（IPv4）
- `IsPhone(s string) bool` - 验证手机号格式（中国大陆）
- `IsNumeric(s string) bool` - 验证是否为数字字符串
- `IsAlpha(s string) bool` - 验证是否为字母字符串
- `IsAlphaNumeric(s string) bool` - 验证是否为字母数字字符串
- `IsEmpty(s string) bool` - 检查字符串是否为空
- `IsBlank(s string) bool` - 检查字符串是否为空白（空或只包含空白字符）

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/str"
)

func main() {
	// 基础操作
	fmt.Println(str.Trim("  hello  "))              // "hello"
	fmt.Println(str.Substring("Hello, World", 0, 5)) // "Hello"
	fmt.Println(str.Reverse("hello"))                // "olleh"
	
	// 格式化
	fmt.Println(str.CamelCase("hello_world"))       // "helloWorld"
	fmt.Println(str.SnakeCase("HelloWorld"))         // "hello_world"
	fmt.Println(str.KebabCase("HelloWorld"))         // "hello-world"
	
	// 验证
	fmt.Println(str.IsEmail("test@example.com"))     // true
	fmt.Println(str.IsURL("https://example.com"))    // true
	fmt.Println(str.IsPhone("13800138000"))          // true
}
```

## 运行测试

```bash
go test ./str -v
```
