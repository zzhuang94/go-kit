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
强大的 HTTP 请求工具和 IP 地址处理，支持 GET、POST、PUT、DELETE、PATCH 等请求方法，JSON 自动序列化/反序列化，文件上传下载，以及完整的响应信息获取。

## 📋 函数列表

### 📁 文件操作 (file)
- `file.Exists(path)` - 检查文件或目录是否存在
- `file.ReadAll(path)` - 读取整个文件内容，返回字节数组
- `file.ReadString(path)` - 读取整个文件内容，返回字符串
- `file.ReadLines(path)` - 读取文件的所有行，返回字符串切片
- `file.Write(path, data)` - 写入字节数组到文件（覆盖模式）
- `file.WriteString(path, content)` - 写入字符串到文件（覆盖模式）
- `file.WriteLines(path, lines)` - 写入多行字符串到文件（覆盖模式）
- `file.Append(path, data)` - 追加字节数组到文件末尾
- `file.AppendString(path, content)` - 追加字符串到文件末尾
- `file.AppendLines(path, lines)` - 追加多行字符串到文件末尾
- `file.Copy(src, dst)` - 复制文件到目标位置
- `file.Move(src, dst)` - 移动/重命名文件
- `file.Remove(path)` - 删除文件
- `file.Create(path)` - 创建空文件
- `file.Touch(path)` - 更新文件访问和修改时间为当前时间
- `file.GetSize(path)` - 获取文件大小（字节）
- `file.GetModTime(path)` - 获取文件修改时间
- `file.IsReadable(path)` - 检查文件是否可读
- `file.IsWritable(path)` - 检查文件是否可写
- `file.IsDir(path)` - 判断指定路径是否为目录
- `file.EnsureDir(dir)` - 确保目录存在，如果不存在则创建（包括所有父目录）
- `file.RemoveDir(dir)` - 删除目录及其所有内容（递归删除）
- `file.CleanDir(dir)` - 清空目录内容但保留目录本身
- `file.MoveDir(src, dst)` - 移动/重命名目录
- `file.ListDir(dir)` - 列出目录下的所有文件和子目录名称
- `file.ListFiles(dir)` - 列出目录下的所有文件名称（不包括子目录）
- `file.ListDirs(dir)` - 列出目录下的所有子目录名称（不包括文件）
- `file.CopyDir(src, dst)` - 复制目录及其所有内容到目标位置
- `file.GetDirSize(dir)` - 获取目录的总大小（字节），递归计算所有文件
- `file.GetDirModTime(dir)` - 获取目录修改时间
- `file.IsDirEmpty(dir)` - 检查目录是否为空
- `file.GetDirCount(dir)` - 获取目录下的文件和子目录数量
- `file.WalkDir(root, walkFn)` - 遍历目录，对每个文件或目录执行指定的函数
- `file.Md5(path)` - 计算文件的 MD5 值
- `file.TarDir(srcDir, dstFile)` - 将目录打包为 tar 文件
- `file.TarGzDir(srcDir, dstFile)` - 将目录打包为 tar.gz 压缩文件
- `file.Untar(srcFile, dstDir)` - 解压 tar 文件到指定目录
- `file.UntarGz(srcFile, dstDir)` - 解压 tar.gz 压缩文件到指定目录

### 🔤 字符串处理 (str)
- `str.Trim(s)` - 去除首尾空白字符
- `str.Md5(params)` - 计算参数的 MD5 值
- `str.Uuid()` - 生成 UUID
- `str.Sub(s, start, end)` - 截取子串，支持负数索引
- `str.Reverse(s)` - 反转字符串
- `str.PadLeft(s, length, pad)` - 左填充字符串到指定长度
- `str.PadRight(s, length, pad)` - 右填充字符串到指定长度
- `str.Truncate(s, length)` - 截断字符串到指定长度
- `str.SplitLines(str)` - 按行分割字符串，去除空行和重复行
- `str.ParseAndFormatJson(str)` - 解析并格式化 JSON 字符串
- `str.CamelCase(s)` - 转换为驼峰命名
- `str.SnakeCase(s)` - 转换为蛇形命名
- `str.KebabCase(s)` - 转换为短横线命名
- `str.TitleCase(s)` - 转换为标题格式（每个单词首字母大写）
- `str.Pluralize(s)` - 简单的复数化（基础规则）
- `str.IsEmail(s)` - 验证邮箱格式
- `str.IsURL(s)` - 验证URL格式
- `str.IsIP(s)` - 验证IP地址格式（IPv4）
- `str.IsPhone(s)` - 验证手机号格式（中国大陆）
- `str.IsNumeric(s)` - 验证是否为数字字符串
- `str.IsAlpha(s)` - 验证是否为字母字符串
- `str.IsAlphaNumeric(s)` - 验证是否为字母数字字符串
- `str.IsEmpty(s)` - 检查字符串是否为空
- `str.IsBlank(s)` - 检查字符串是否为空白（空或只包含空白字符）

### 📊 切片操作 (slice)
- `slice.Contains(slice, item)` - 检查切片是否包含指定元素
- `slice.IndexOf(slice, item)` - 查找元素在切片中的索引，不存在返回 -1
- `slice.LastIndexOf(slice, item)` - 从后往前查找元素在切片中的索引，不存在返回 -1
- `slice.Unique(slice)` - 去除重复元素
- `slice.Reverse(slice)` - 反转切片
- `slice.Shuffle(slice)` - 随机打乱切片
- `slice.Chunk(slice, size)` - 将切片分块
- `slice.Flatten(slices)` - 扁平化嵌套切片
- `slice.Intersect(slice1, slice2)` - 求两个切片的交集
- `slice.Union(slice1, slice2)` - 求两个切片的并集
- `slice.Diff(slice1, slice2)` - 求两个切片的差集（slice1 中有但 slice2 中没有的元素）
- `slice.Remove(slice, item)` - 删除切片中所有匹配的元素
- `slice.RemoveAt(slice, index)` - 删除指定索引的元素
- `slice.Insert(slice, index, item)` - 在指定索引插入元素
- `slice.First(slice)` - 获取第一个元素，如果切片为空返回零值
- `slice.Last(slice)` - 获取最后一个元素，如果切片为空返回零值
- `slice.Take(slice, n)` - 取前N个元素
- `slice.Drop(slice, n)` - 跳过前N个元素
- `slice.Filter(slice, fn)` - 过滤切片，保留满足条件的元素
- `slice.Find(slice, fn)` - 查找第一个满足条件的元素，不存在返回零值
- `slice.FindIndex(slice, fn)` - 查找第一个满足条件的元素索引，不存在返回 -1
- `slice.FindLast(slice, fn)` - 从后往前查找第一个满足条件的元素，不存在返回零值
- `slice.FindLastIndex(slice, fn)` - 从后往前查找第一个满足条件的元素索引，不存在返回 -1
- `slice.Every(slice, fn)` - 检查是否所有元素都满足条件
- `slice.Some(slice, fn)` - 检查是否至少有一个元素满足条件
- `slice.Count(slice, fn)` - 统计满足条件的元素个数
- `slice.GroupBy(slice, fn)` - 按指定函数分组
- `slice.Map(slice, fn)` - 映射转换切片
- `slice.Reduce(slice, initial, fn)` - 归约切片
- `slice.FlatMap(slice, fn)` - 扁平映射
- `slice.Partition(slice, fn)` - 将切片分为两部分：满足条件的和不满足条件的

### 🔢 算法与数据结构 (algo)
- `algo.NewStack[T]()` - 创建新栈
- `stack.Push(item)` - 入栈
- `stack.Pop()` - 出栈
- `stack.Peek()` - 查看栈顶元素
- `stack.Size()` - 返回栈大小
- `stack.IsEmpty()` - 检查栈是否为空
- `stack.Clear()` - 清空栈
- `algo.NewQueue[T]()` - 创建新队列
- `queue.Enqueue(item)` - 入队
- `queue.Dequeue()` - 出队
- `queue.Peek()` - 查看队首元素
- `queue.Size()` - 返回队列大小
- `queue.IsEmpty()` - 检查队列是否为空
- `queue.Clear()` - 清空队列
- `algo.NewSet[T]()` - 创建新集合
- `set.Add(item)` - 添加元素
- `set.Remove(item)` - 删除元素
- `set.Contains(item)` - 检查是否包含元素
- `set.Size()` - 返回集合大小
- `set.IsEmpty()` - 检查集合是否为空
- `set.Clear()` - 清空集合
- `set.ToSlice()` - 转换为切片
- `set.Union(other)` - 求并集
- `set.Intersect(other)` - 求交集
- `set.Diff(other)` - 求差集（s 中有但 other 中没有的元素）
- `algo.NewLRU[K, V](capacity)` - 创建LRU缓存
- `lru.Get(key)` - 获取值
- `lru.Set(key, value)` - 设置值
- `lru.Remove(key)` - 删除键值对
- `lru.Clear()` - 清空缓存
- `lru.Size()` - 返回当前大小
- `lru.Cap()` - 返回容量
- `algo.Sort(slice)` - 对切片进行排序（使用快速排序）
- `algo.SortFunc(slice, less)` - 使用自定义比较函数对切片进行排序
- `algo.SortDesc(slice)` - 对切片进行降序排序
- `algo.IsSorted(slice)` - 检查切片是否已排序
- `algo.IsSortedFunc(slice, less)` - 使用自定义比较函数检查切片是否已排序
- `algo.BinarySearch(slice, target)` - 二分查找，返回目标值的索引，如果不存在返回 -1
- `algo.BinarySearchFunc(slice, target, cmp)` - 使用自定义比较函数进行二分查找
- `algo.LinearSearch(slice, target)` - 线性查找，返回目标值的索引，如果不存在返回 -1
- `algo.FindFirst(slice, fn)` - 查找第一个满足条件的元素索引，不存在返回 -1
- `algo.FindLast(slice, fn)` - 查找最后一个满足条件的元素索引，不存在返回 -1
- `algo.LowerBound(slice, target)` - 查找第一个大于等于目标值的元素索引
- `algo.UpperBound(slice, target)` - 查找第一个大于目标值的元素索引

### ⏰ 时间处理 (time)
- `time.Format(t, layout)` - 格式化时间为字符串
- `time.FormatDate(t)` - 格式化日期（YYYY-MM-DD）
- `time.FormatTime(t)` - 格式化时间（HH:mm:ss）
- `time.FormatDateTime(t)` - 格式化日期时间（YYYY-MM-DD HH:mm:ss）
- `time.FormatTimestamp(t)` - 格式化时间戳
- `time.Parse(layout, value)` - 解析时间字符串
- `time.ParseDate(value)` - 解析日期字符串（YYYY-MM-DD）
- `time.ParseTime(value)` - 解析时间字符串（HH:mm:ss）
- `time.ParseDateTime(value)` - 解析日期时间字符串（YYYY-MM-DD HH:mm:ss）
- `time.AddDays(t, days)` - 添加天数
- `time.AddHours(t, hours)` - 添加小时
- `time.AddMinutes(t, minutes)` - 添加分钟
- `time.DiffDays(t1, t2)` - 计算天数差
- `time.DiffHours(t1, t2)` - 计算小时差
- `time.DiffMinutes(t1, t2)` - 计算分钟差
- `time.IsToday(t)` - 判断是否为今天
- `time.IsYesterday(t)` - 判断是否为昨天
- `time.IsTomorrow(t)` - 判断是否为明天
- `time.StartOfDay(t)` - 获取一天的开始时间
- `time.EndOfDay(t)` - 获取一天的结束时间
- `time.StartOfWeek(t)` - 获取一周的开始时间（周一）
- `time.StartOfMonth(t)` - 获取一月的开始时间
- `time.StartOfYear(t)` - 获取一年的开始时间

### 🔐 加密工具 (crypto)
- `crypto.SHA1(data)` - 计算字节数组的 SHA1 值
- `crypto.SHA256(data)` - 计算字节数组的 SHA256 值
- `crypto.SHA512(data)` - 计算字节数组的 SHA512 值
- `crypto.SHA1String(s)` - 计算字符串的 SHA1 值
- `crypto.SHA256String(s)` - 计算字符串的 SHA256 值
- `crypto.SHA512String(s)` - 计算字符串的 SHA512 值
- `crypto.SHA1File(path)` - 计算文件的 SHA1 值
- `crypto.SHA256File(path)` - 计算文件的 SHA256 值
- `crypto.SHA512File(path)` - 计算文件的 SHA512 值
- `crypto.AESEncrypt(key, plaintext)` - AES加密
- `crypto.AESDecrypt(key, ciphertext)` - AES解密
- `crypto.AESEncryptString(key, plaintext)` - AES加密（字符串）
- `crypto.AESDecryptString(key, ciphertext)` - AES解密（字符串）
- `crypto.Base64Encode(data)` - Base64编码
- `crypto.Base64Decode(data)` - Base64解码
- `crypto.Base64EncodeString(s)` - Base64编码（字符串）
- `crypto.Base64DecodeString(s)` - Base64解码（字符串）

### 🛠️ 通用工具库 (lib)
- `lib.Choice(slice)` - 从切片中随机选择一个元素
- `lib.Shuffle(slice)` - 随机打乱切片
- `lib.RunCmd(cmd, timeout...)` - 执行 Shell 命令，支持超时设置
- `logCfg.InitLogrus()` - 初始化全局 logrus 日志
- `logCfg.BuildLogger()` - 创建新的 logrus Logger 实例
- `lib.GetFormatter()` - 获取自定义日志格式化器
- `lib.TryLock(c, key, timeout)` - 尝试获取分布式锁
- `lock.Release()` - 释放锁
- `lib.TryCheckIn(c, key, limit, timeout)` - 尝试进入限流器
- `limiter.Release()` - 释放限流器资源
- `lib.NewElection(key, val, cmdable)` - 创建选举实例
- `election.IsMaster()` - 检查当前实例是否为主节点
- `election.GetMaster()` - 获取当前主节点的值
- `election.Release()` - 释放主节点身份
- `election.Stop()` - 停止选举并释放资源

### 💾 数据库操作 (db)
- `mysqlCfg.ConnGorm()` - 使用 GORM 连接 MySQL
- `mysqlCfg.ConnXorm()` - 使用 XORM 连接 MySQL
- `db.ConnRedis(cfg)` - 连接单机 Redis
- `db.ConnRedisCluster(cfg)` - 连接 Redis 集群

### 🌐 网络操作 (net)
- `net.Get(url, timeout)` - 发送 GET 请求
- `net.GetWithHeaders(url, timeout, headers)` - 发送带请求头的 GET 请求
- `net.Post(url, data, headers, timeout)` - 发送 POST 请求
- `net.Put(url, data, headers, timeout)` - 发送 PUT 请求
- `net.Delete(url, headers, timeout)` - 发送 DELETE 请求
- `net.Patch(url, data, headers, timeout)` - 发送 PATCH 请求
- `net.Req(method, url, data, headers, timeout)` - 发送通用 HTTP 请求
- `net.ReqFull(method, url, data, headers, timeout)` - 发送请求并返回完整响应信息
- `net.PostJson(url, data, headers, timeout)` - 发送 JSON POST 请求
- `net.GetJsonParse(url, result, timeout)` - 发送 GET 请求并解析 JSON 响应
- `net.PostJsonParse(url, data, result, headers, timeout)` - 发送 JSON POST 请求并解析 JSON 响应
- `net.UploadFile(url, path, name, headers, timeout)` - 上传文件
- `net.DownloadFile(url, path, timeout)` - 下载文件到本地
- `net.ClientIp(req)` - 获取客户端 IP 地址
- `net.LocalIp()` - 获取本机 IP 地址

## 🚀 快速开始

### 安装

安装最新版本：
```bash
go get github.com/zzhuang94/go-kit@latest
```

安装指定版本（推荐）：
```bash
go get github.com/zzhuang94/go-kit@v1.0.0
```

如果遇到版本问题，可以在项目的 `go.mod` 文件中直接指定版本：
```go
require (
    github.com/zzhuang94/go-kit v1.0.0
)
```
然后运行 `go mod tidy`。

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
