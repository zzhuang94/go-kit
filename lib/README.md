# 通用工具库 (lib)

提供常用的通用工具函数，包括 Shell 命令执行、日志管理、分布式锁、限流器和分布式选举等功能。

## Shell 命令执行 (shell)

- `RunCmd(cmd string, timeout ...int) (string, int, error)` - 执行 Shell 命令，支持超时设置
  - `cmd`: 要执行的命令
  - `timeout`: 可选的超时时间（秒），不设置则无超时限制
  - 返回: 命令输出、退出码和错误信息

## 日志管理 (log)

基于 logrus 的日志管理工具，支持日志轮转和自定义格式化。

- `LogCfg` - 日志配置结构体
  - `Path string` - 日志文件路径
  - `Level string` - 日志级别（debug, info, warn, error）
  - `KeepDays time.Duration` - 日志保留天数（默认 3 天）
- `LogCfg.InitLogrus()` - 初始化全局 logrus 日志
- `LogCfg.BuildLogger() *logrus.Logger` - 创建新的 logrus Logger 实例
- `LogCfg.ParseLevel() logrus.Level` - 解析日志级别
- `LogCfg.BuildLogWriter() io.Writer` - 构建日志写入器（支持日志轮转）
- `GetFormatter() logrus.Formatter` - 获取自定义日志格式化器

## 分布式锁 (lock)

基于限流器实现的分布式锁，支持 Redis、etcd 和本地模式。

- `TryLockWithRedis(c redis.Cmdable, key string, timeout time.Duration) (*Lock, error)` - 使用 Redis 尝试获取分布式锁
- `TryLockWithEtcd(c *clientv3.Client, key string, timeout time.Duration) (*Lock, error)` - 使用 etcd 尝试获取分布式锁
- `TryLockLocal(key string, timeout time.Duration) (*Lock, error)` - 使用本地模式尝试获取分布式锁
- `Lock.Release()` - 释放锁

## 限流器 (limiter)

支持 Redis、etcd 和本地三种模式的限流器，用于控制并发访问数量。

- `TryCheckInWithRedis(c redis.Cmdable, key string, limit int, timeout time.Duration) (*Limiter, error)` - 使用 Redis 尝试进入限流器
- `TryCheckInWithEtcd(c *clientv3.Client, key string, limit int, timeout time.Duration) (*Limiter, error)` - 使用 etcd 尝试进入限流器
- `TryCheckInLocal(key string, limit int, timeout time.Duration) (*Limiter, error)` - 使用本地模式尝试进入限流器
- `Limiter.Release()` - 释放限流器资源

参数说明：
- `key`: 限流器的键名
- `limit`: 最大并发数
- `timeout`: 获取限流器资源的超时时间

## 分布式选举 (election)

基于 Redis 或 etcd 的分布式选举实现，用于在多个实例中选择主节点。

- `NewElectionWithRedis(cmdable redis.Cmdable, key, val string) *Election` - 使用 Redis 创建选举实例
- `NewElectionWithEtcd(client *clientv3.Client, key, val string) *Election` - 使用 etcd 创建选举实例
- `Election.IsMaster() bool` - 检查当前实例是否为主节点
- `Election.GetMaster() string` - 获取当前主节点的值
- `Election.Release()` - 释放主节点身份
- `Election.Stop()` - 停止选举并释放资源

参数说明：
- `key`: 选举的键名
- `val`: 当前实例的标识值

## 存储接口 (storage)

内部存储接口实现，支持本地内存、Redis 和 etcd 三种存储方式。

- `NewLocalStor(key string) *LocalStor` - 创建本地内存存储
- `NewRedisStor(c redis.Cmdable, key string) *RedisStor` - 创建 Redis 存储
- `NewEtcdStor(client *clientv3.Client, key string) *EtcdStor` - 创建 etcd 存储

这些存储接口主要用于支持分布式锁、限流器和选举功能。

## 使用示例

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/zzhuang94/go-kit/lib"
	redis "github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	// Shell 命令执行
	output, code, err := lib.RunCmd("echo hello", 5)
	if err == nil {
		fmt.Println("Output:", output, "Exit code:", code)
	}
	
	// 日志管理
	cfg := &lib.LogCfg{
		Path:     "./logs/app.log",
		Level:    "info",
		KeepDays: 7 * 24 * time.Hour,
	}
	cfg.InitLogrus()
	logrus.Info("Application started")
	
	// 或者创建独立的 logger
	logger := cfg.BuildLogger()
	logger.Info("Using custom logger")
	
	// Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	// 分布式锁（Redis）
	lock, err := lib.TryLockWithRedis(rdb, "my_lock", 10*time.Second)
	if err == nil {
		defer lock.Release()
		fmt.Println("Lock acquired!")
		// 执行需要加锁的操作
	}
	
	// 分布式锁（本地模式）
	lock, err = lib.TryLockLocal("local_lock", 5*time.Second)
	if err == nil {
		defer lock.Release()
		fmt.Println("Local lock acquired!")
	}
	
	// 限流器（Redis）
	limiter, err := lib.TryCheckInWithRedis(rdb, "api_limit", 100, 5*time.Second)
	if err == nil {
		defer limiter.Release()
		fmt.Println("Rate limiter check-in successful!")
		// 执行需要限流的操作
	}
	
	// 限流器（本地模式）
	limiter, err = lib.TryCheckInLocal("local_limit", 10, 3*time.Second)
	if err == nil {
		defer limiter.Release()
		fmt.Println("Local rate limiter check-in successful!")
	}
	
	// 分布式选举（Redis）
	election := lib.NewElectionWithRedis(rdb, "master_key", "node1")
	if election.IsMaster() {
		fmt.Println("I am the master!")
	}
	defer election.Stop()
	
	// 分布式选举（etcd）
	etcdClient, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	election = lib.NewElectionWithEtcd(etcdClient, "master_key", "node2")
	if election.IsMaster() {
		fmt.Println("I am the master!")
	}
	master := election.GetMaster()
	fmt.Println("Current master:", master)
	defer election.Stop()
}
```

## 运行测试

```bash
go test ./lib -v
```
