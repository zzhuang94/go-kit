# General Utilities (lib)

Provides common general utility functions including Shell command execution, log management, distributed locks, rate limiters, and distributed election.

## Shell Command Execution (shell)

- `RunCmd(cmd string, timeout ...int) (string, int, error)` - Execute Shell command with optional timeout
  - `cmd`: Command to execute
  - `timeout`: Optional timeout in seconds, no timeout if not set
  - Returns: Command output, exit code, and error

## Log Management (log)

Log management utilities based on logrus with log rotation and custom formatting support.

- `LogCfg` - Log configuration struct
  - `Path string` - Log file path
  - `Level string` - Log level (debug, info, warn, error)
  - `KeepDays time.Duration` - Log retention days (default 3 days)
- `LogCfg.InitLogrus()` - Initialize global logrus logger
- `LogCfg.BuildLogger() *logrus.Logger` - Create new logrus Logger instance
- `LogCfg.ParseLevel() logrus.Level` - Parse log level
- `LogCfg.BuildLogWriter() io.Writer` - Build log writer (supports log rotation)
- `GetFormatter() logrus.Formatter` - Get custom log formatter

## Distributed Lock (lock)

Distributed lock implementation based on rate limiter, supporting Redis, etcd, and local modes.

- `TryLockWithRedis(c redis.Cmdable, key string, timeout time.Duration) (*Lock, error)` - Try to acquire distributed lock using Redis
- `TryLockWithEtcd(c *clientv3.Client, key string, timeout time.Duration) (*Lock, error)` - Try to acquire distributed lock using etcd
- `TryLockLocal(key string, timeout time.Duration) (*Lock, error)` - Try to acquire distributed lock using local mode
- `Lock.Release()` - Release lock

## Rate Limiter (limiter)

Rate limiter supporting Redis, etcd, and local modes for controlling concurrent access.

- `TryCheckInWithRedis(c redis.Cmdable, key string, limit int, timeout time.Duration) (*Limiter, error)` - Try to check in to rate limiter using Redis
- `TryCheckInWithEtcd(c *clientv3.Client, key string, limit int, timeout time.Duration) (*Limiter, error)` - Try to check in to rate limiter using etcd
- `TryCheckInLocal(key string, limit int, timeout time.Duration) (*Limiter, error)` - Try to check in to rate limiter using local mode
- `Limiter.Release()` - Release rate limiter resource

Parameters:
- `key`: Rate limiter key name
- `limit`: Maximum concurrent count
- `timeout`: Timeout for acquiring rate limiter resource

## Distributed Election (election)

Distributed election implementation based on Redis or etcd for selecting master node among multiple instances.

- `NewElectionWithRedis(cmdable redis.Cmdable, key, val string) *Election` - Create election instance using Redis
- `NewElectionWithEtcd(client *clientv3.Client, key, val string) *Election` - Create election instance using etcd
- `Election.IsMaster() bool` - Check if current instance is master
- `Election.GetMaster() string` - Get current master node value
- `Election.Release()` - Release master status
- `Election.Stop()` - Stop election and release resources

Parameters:
- `key`: Election key name
- `val`: Current instance identifier value

## Storage Interface (storage)

Internal storage interface implementations supporting local memory, Redis, and etcd storage.

- `NewLocalStor(key string) *LocalStor` - Create local memory storage
- `NewRedisStor(c redis.Cmdable, key string) *RedisStor` - Create Redis storage
- `NewEtcdStor(client *clientv3.Client, key string) *EtcdStor` - Create etcd storage

These storage interfaces are primarily used to support distributed locks, rate limiters, and election features.

## Usage Examples

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
	// Shell command execution
	output, code, err := lib.RunCmd("echo hello", 5)
	if err == nil {
		fmt.Println("Output:", output, "Exit code:", code)
	}
	
	// Log management
	cfg := &lib.LogCfg{
		Path:     "./logs/app.log",
		Level:    "info",
		KeepDays: 7 * 24 * time.Hour,
	}
	cfg.InitLogrus()
	logrus.Info("Application started")
	
	// Or create independent logger
	logger := cfg.BuildLogger()
	logger.Info("Using custom logger")
	
	// Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	// Distributed lock (Redis)
	lock, err := lib.TryLockWithRedis(rdb, "my_lock", 10*time.Second)
	if err == nil {
		defer lock.Release()
		fmt.Println("Lock acquired!")
		// Execute operations requiring lock
	}
	
	// Distributed lock (local mode)
	lock, err = lib.TryLockLocal("local_lock", 5*time.Second)
	if err == nil {
		defer lock.Release()
		fmt.Println("Local lock acquired!")
	}
	
	// Rate limiter (Redis)
	limiter, err := lib.TryCheckInWithRedis(rdb, "api_limit", 100, 5*time.Second)
	if err == nil {
		defer limiter.Release()
		fmt.Println("Rate limiter check-in successful!")
		// Execute operations requiring rate limiting
	}
	
	// Rate limiter (local mode)
	limiter, err = lib.TryCheckInLocal("local_limit", 10, 3*time.Second)
	if err == nil {
		defer limiter.Release()
		fmt.Println("Local rate limiter check-in successful!")
	}
	
	// Distributed election (Redis)
	election := lib.NewElectionWithRedis(rdb, "master_key", "node1")
	if election.IsMaster() {
		fmt.Println("I am the master!")
	}
	defer election.Stop()
	
	// Distributed election (etcd)
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

## Running Tests

```bash
go test ./lib -v
```
