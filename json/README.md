# JSON 工具包

提供常用的JSON操作工具函数。

## 序列化和反序列化

- `Marshal(v interface{}) ([]byte, error)` - JSON序列化
- `Unmarshal(data []byte, v interface{}) error` - JSON反序列化
- `MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)` - JSON序列化（格式化）
- `PrettyPrint(v interface{}) error` - 美化打印JSON

## JSON操作

- `Get(data []byte, key string) (interface{}, error)` - 获取JSON字段值
- `Set(data []byte, key string, value interface{}) ([]byte, error)` - 设置JSON字段值
- `Merge(data1, data2 []byte) ([]byte, error)` - 合并JSON对象

## 文件操作

- `ReadFile(path string, v interface{}) error` - 从文件读取JSON
- `WriteFile(path string, v interface{}) error` - 写入JSON到文件

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/json"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	
	// 序列化
	p := Person{Name: "Alice", Age: 30}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	
	// 反序列化
	var p2 Person
	json.Unmarshal(data, &p2)
	fmt.Println(p2)
	
	// 美化打印
	json.PrettyPrint(p)
	
	// 获取字段值
	value, _ := json.Get(data, "name")
	fmt.Println(value)
	
	// 设置字段值
	newData, _ := json.Set(data, "age", 31)
	fmt.Println(string(newData))
}
```

## 运行测试

```bash
go test ./json -v
```
