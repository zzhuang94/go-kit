# 🚀 go-kit

> 一个功能丰富、开箱即用的 Go 语言工具库集合，让开发更高效、代码更优雅

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zzhuang94/go-kit?style=flat-square)](https://goreportcard.com/report/github.com/zzhuang94/go-kit)

**go-kit** 是一个精心设计的 Go 语言工具库集合，旨在为开发者提供常用功能的简洁实现。无论是文件操作、字符串处理、数据结构，还是时间处理、加密解密，这里都有你需要的工具函数。所有函数都经过精心设计，支持泛型，提供完整的错误处理，并附带详细的中英文文档。

## ✨ 特性

- 🎯 **开箱即用** - 无需复杂配置，导入即用
- 🔧 **功能丰富** - 涵盖文件、字符串、数据结构、时间、加密等 9+ 个模块
- 🚀 **高性能** - 基于 Go 标准库，性能优异
- 📚 **文档完善** - 每个函数都有中英文注释和详细文档
- 🧪 **测试完备** - 完整的单元测试覆盖
- 🔒 **类型安全** - 充分利用 Go 泛型，保证类型安全
- 🌍 **国际化** - 提供中英文双语文档

## 📦 功能模块

### 📁 文件操作 (file)
强大的文件和目录操作工具，支持文件读写、复制、移动、删除，目录创建、遍历、复制，以及 tar/tar.gz 压缩解压和 MD5 哈希计算。

### 🔤 字符串处理 (str)
丰富的字符串操作函数，包括基础操作（截取、替换、反转等）、命名格式转换（驼峰、蛇形、短横线等）和常用验证（邮箱、URL、IP、手机号等）。

### 📊 切片操作 (slice)
强大的切片操作工具，支持去重、反转、分块、过滤、查找、映射、归约等函数式编程操作，充分利用 Go 泛型特性。

### 🔢 算法与数据结构 (algo)
常用数据结构和算法的泛型实现，包括栈（Stack）、队列（Queue）、集合（Set）、LRU 缓存，以及排序算法（快速排序）和搜索算法（二分查找、线性查找等），性能优异，API 简洁。

### ⏰ 时间处理 (time)
便捷的时间操作工具，支持格式化、解析、计算（加减、差值、判断等），让时间处理变得简单直观。

### 🔐 加密工具 (crypto)
完整的加密工具集，包括 SHA 系列哈希（SHA1、SHA256、SHA512）、AES 加密解密和 Base64 编码解码。

### 🛠️ 通用工具库 (lib)
实用的通用工具集合，包括随机数操作（随机选择、打乱）、Shell 命令执行、日志管理（基于 logrus，支持日志轮转）、分布式锁（基于 Redis）、限流器（支持 Redis 和本地模式）和分布式选举（基于 Redis）。

### 💾 数据库操作 (db)
数据库连接工具，支持 MySQL（GORM、XORM）和 Redis（单机、集群）连接。

### 🌐 网络操作 (net)
HTTP 请求工具和 IP 地址处理，简化网络编程。

## 🚀 快速开始

### 安装

```bash
go get github.com/zzhuang94/go-kit
```

### 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/file"
	"github.com/zzhuang94/go-kit/str"
	"github.com/zzhuang94/go-kit/slice"
	"github.com/zzhuang94/go-kit/time"
)

func main() {
	// 文件操作
	file.WriteString("test.txt", "Hello, World!")
	content, _ := file.ReadString("test.txt")
	fmt.Println(content)
	
	// 字符串处理
	fmt.Println(str.CamelCase("hello_world"))  // "helloWorld"
	fmt.Println(str.IsEmail("test@example.com")) // true
	
	// 切片操作
	numbers := []int{1, 2, 3, 2, 4}
	unique := slice.Unique(numbers)
	fmt.Println(unique) // [1, 2, 3, 4]
	
	// 时间处理
	now := time.Now()
	fmt.Println(time.FormatDateTime(now)) // "2024-01-15 10:30:45"
}
```

## 📚 文档

每个模块都有详细的文档和使用示例，请查看各子目录下的 README.md 文件：

- [文件操作 (file)](./file/README.md)
- [字符串处理 (str)](./str/README.md)
- [切片操作 (slice)](./slice/README.md)
- [算法与数据结构 (algo)](./algo/README.md)
- [时间处理 (time)](./time/README.md)
- [加密工具 (crypto)](./crypto/README.md)
- [通用工具库 (lib)](./lib/README.md)
- [数据库操作 (db)](./db/README.md)
- [网络操作 (net)](./net/README.md)

## 🧪 运行测试

```bash
# 运行所有测试（包括单元测试和基准测试）
make test

# 或使用一键测试脚本
# Windows
test.bat

# Unix/Linux/Mac
./test.sh

# 或使用 Go 命令
go test ./...

# 运行测试并显示覆盖率
go test ./... -cover

# 只运行基准测试
make test-bench
# 或
go test ./... -bench=. -benchmem
```

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

本项目采用 [MIT License](LICENSE) 许可证。

---

⭐ 如果这个项目对你有帮助，请给它一个 Star！
