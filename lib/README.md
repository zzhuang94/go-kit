# 通用工具库 (lib)

提供常用的通用工具函数，包括随机数操作、Shell 命令执行、日志管理、分布式锁、限流器和分布式选举等功能。

## 随机数操作 (rand)

- `Choice[T any](slice []T) (T, bool)` - 从切片中随机选择一个元素
- `Shuffle[T any](slice []T) []T` - 随机打乱切片

## Shell 命令执行 (shell)

- `RunCmd(cmd string, timeout ...int) (string, int, error)` - 执行 Shell 命令，支持超时设置

## 日志管理 (log)

基于 logrus 的日志管理工具，支持日志轮转和自定义格式化。

- `LogCfg` - 日志配置结构体
  - `Path string` - 日志文件路径
  - `Level string` - 日志级别（debug, info, warn, error）
  - `KeepDays time.Duration` - 日志保留天数
- `LogCfg.InitLogrus()` - 初始化全局 logrus 日志
- `LogCfg.BuildLogger() *logrus.Logger` - 创建新的 logrus Logger 实例
- `GetFormatter() logrus.Formatter` - 获取自定义日志格式化器

## 分布式锁 (lock)

基于 Redis 的分布式锁实现。

- `TryLock(c redis.Cmdable, key string, timeout time.Duration) (*lock, error)` - 尝试获取分布式锁
- `lock.Release()` - 释放锁

## 限流器 (limiter)

支持 Redis 和本地两种模式的限流器。

- `TryCheckIn(c redis.Cmdable, key string, limit int, timeout time.Duration) (*limiter, error)` - 尝试进入限流器
- `limiter.Release()` - 释放限流器资源

## 分布式选举 (election)

基于 Redis 的分布式选举实现，用于在多个实例中选择主节点。

- `NewElection(key, val string, cmdable redis.Cmdable) *Election` - 创建选举实例
- `Election.IsMaster() bool` - 检查当前实例是否为主节点
- `Election.GetMaster() string` - 获取当前主节点的值
- `Election.Release()` - 释放主节点身份
- `Election.Stop()` - 停止选举并释放资源

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/lib"
	redis "github.com/go-redis/redis/v8"
)

func main() {
	// 随机数操作
	numbers := []int{1, 2, 3, 4, 5}
	value, _ := lib.Choice(numbers)
	fmt.Println("Random choice:", value)
	
	shuffled := lib.Shuffle(numbers)
	fmt.Println("Shuffled:", shuffled)
	
	// Shell 命令执行
	output, code, err := lib.RunCmd("echo hello", 5)
	if err == nil {
		fmt.Println("Output:", output)
	}
	
	// 日志管理
	cfg := &lib.LogCfg{
		Path:     "./logs/app.log",
		Level:    "info",
		KeepDays: 7,
	}
	cfg.InitLogrus()
	
	// 分布式锁
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	lock, err := lib.TryLock(rdb, "my_lock", 10*time.Second)
	if err == nil {
		defer lock.Release()
		// 执行需要加锁的操作
	}
	
	// 限流器
	limiter, err := lib.TryCheckIn(rdb, "api_limit", 100, 5*time.Second)
	if err == nil {
		defer limiter.Release()
		// 执行需要限流的操作
	}
	
	// 分布式选举
	election := lib.NewElection("master_key", "node1", rdb)
	if election.IsMaster() {
		fmt.Println("I am the master!")
	}
	defer election.Stop()
}
```

## 运行测试

```bash
go test ./lib -v
```
