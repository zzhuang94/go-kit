# go-kit

Go 语言常用工具库集合，提供文件操作、字符串处理、数据结构、时间处理、加密等常用功能。

## 功能模块

### 文件操作 (file)
- 文件读写、复制、移动、删除
- 目录操作（创建、删除、列表、复制）
- 文件压缩解压（tar、tar.gz）
- MD5 哈希计算

### 字符串处理 (str)
- 字符串基础操作（截取、替换、反转等）
- 命名格式转换（驼峰、蛇形、短横线等）
- 字符串验证（邮箱、URL、IP、手机号等）

### 切片操作 (slice)
- 切片基础操作（去重、反转、分块等）
- 过滤和查找
- 映射和归约转换

### 数据结构 (structs)
- 栈（Stack）
- 队列（Queue）
- 集合（Set）
- LRU 缓存

### 时间处理 (time)
- 时间格式化
- 时间解析
- 时间计算（加减、差值、判断等）

### 加密工具 (crypto)
- SHA 系列哈希（SHA1、SHA256、SHA512）
- AES 加密解密
- Base64 编码解码

### 类型转换 (convert)
- 字符串与数字互转
- 字符串与布尔值互转
- 字节数组与字符串互转

### JSON 操作 (json)
- JSON 序列化和反序列化
- JSON 美化打印
- JSON 字段操作（获取、设置、合并）

### 随机数生成 (rand)
- 随机整数、浮点数
- 随机字符串生成
- 随机选择元素

### 数据库操作 (db)
- MySQL 连接（GORM、XORM）
- Redis 连接

### 网络操作 (net)
- HTTP 请求（GET、POST）
- IP 地址处理

## 安装

```bash
go get github.com/zzhuang94/go-kit
```

## 快速开始

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/file"
	"github.com/zzhuang94/go-kit/str"
	"github.com/zzhuang94/go-kit/slice"
)

func main() {
	// 文件操作
	file.WriteString("test.txt", "Hello, World!")
	content, _ := file.ReadString("test.txt")
	fmt.Println(content)
	
	// 字符串处理
	fmt.Println(str.CamelCase("hello_world")) // "helloWorld"
	
	// 切片操作
	numbers := []int{1, 2, 3, 2, 4}
	unique := slice.Unique(numbers)
	fmt.Println(unique) // [1, 2, 3, 4]
}
```

## 文档

每个模块都有详细的文档，请查看各子目录下的 README.md 文件。

## 许可证

MIT License
