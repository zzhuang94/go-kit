# General Utilities (lib)

Provides common general utility functions including random operations, Shell command execution, log management, distributed locks, rate limiters, and distributed election.

## Random Operations (rand)

- `Choice[T any](slice []T) (T, bool)` - Randomly select an element from slice
- `Shuffle[T any](slice []T) []T` - Randomly shuffle slice

## Shell Command Execution (shell)

- `RunCmd(cmd string, timeout ...int) (string, int, error)` - Execute Shell command with optional timeout

## Log Management (log)

Log management utilities based on logrus with log rotation and custom formatting support.

- `LogCfg` - Log configuration struct
  - `Path string` - Log file path
  - `Level string` - Log level (debug, info, warn, error)
  - `KeepDays time.Duration` - Log retention days
- `LogCfg.InitLogrus()` - Initialize global logrus logger
- `LogCfg.BuildLogger() *logrus.Logger` - Create new logrus Logger instance
- `GetFormatter() logrus.Formatter` - Get custom log formatter

## Distributed Lock (lock)

Redis-based distributed lock implementation.

- `TryLock(c redis.Cmdable, key string, timeout time.Duration) (*lock, error)` - Try to acquire distributed lock
- `lock.Release()` - Release lock

## Rate Limiter (limiter)

Rate limiter supporting both Redis and local modes.

- `TryCheckIn(c redis.Cmdable, key string, limit int, timeout time.Duration) (*limiter, error)` - Try to check in to rate limiter
- `limiter.Release()` - Release rate limiter resource

## Distributed Election (election)

Redis-based distributed election implementation for selecting master node among multiple instances.

- `NewElection(key, val string, cmdable redis.Cmdable) *Election` - Create election instance
- `Election.IsMaster() bool` - Check if current instance is master
- `Election.GetMaster() string` - Get current master node value
- `Election.Release()` - Release master status
- `Election.Stop()` - Stop election and release resources

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/lib"
	redis "github.com/go-redis/redis/v8"
)

func main() {
	// Random operations
	numbers := []int{1, 2, 3, 4, 5}
	value, _ := lib.Choice(numbers)
	fmt.Println("Random choice:", value)
	
	shuffled := lib.Shuffle(numbers)
	fmt.Println("Shuffled:", shuffled)
	
	// Shell command execution
	output, code, err := lib.RunCmd("echo hello", 5)
	if err == nil {
		fmt.Println("Output:", output)
	}
	
	// Log management
	cfg := &lib.LogCfg{
		Path:     "./logs/app.log",
		Level:    "info",
		KeepDays: 7,
	}
	cfg.InitLogrus()
	
	// Distributed lock
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	lock, err := lib.TryLock(rdb, "my_lock", 10*time.Second)
	if err == nil {
		defer lock.Release()
		// Execute operations requiring lock
	}
	
	// Rate limiter
	limiter, err := lib.TryCheckIn(rdb, "api_limit", 100, 5*time.Second)
	if err == nil {
		defer limiter.Release()
		// Execute operations requiring rate limiting
	}
	
	// Distributed election
	election := lib.NewElection("master_key", "node1", rdb)
	if election.IsMaster() {
		fmt.Println("I am the master!")
	}
	defer election.Stop()
}
```

## Running Tests

```bash
go test ./lib -v
```
