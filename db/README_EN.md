# Database Utilities Package

Provides database connection utility functions.

## MySQL Connection

### GORM Connection

```go
type MysqlCfg struct {
	DSN     string `json:"dsn"`      // Data source name
	MaxIdle int    `json:"max_idle"` // Maximum idle connections
	MaxOpen int    `json:"max_open"` // Maximum open connections
	Log     bool   `json:"log"`      // Enable logging
}

func (c *MysqlCfg) ConnGorm() (*gorm.DB, error) // Connect MySQL using GORM
```

### XORM Connection

```go
func (c *MysqlCfg) ConnXorm() (*xorm.Engine, error) // Connect MySQL using XORM
```

## Redis Connection

### Standalone Redis

```go
type RedisCfg struct {
	Addr   string `json:"addr"`   // Redis address
	Passwd string `json:"passwd"` // Redis password
}

func ConnRedis(c *RedisCfg) redis.Cmdable // Connect standalone Redis
```

### Redis Cluster

```go
type RedisClusterCfg struct {
	Addrs  []string `json:"addrs"`  // Redis cluster addresses
	Passwd string   `json:"passwd"`  // Redis password
}

func ConnRedisCluster(c *RedisClusterCfg) redis.Cmdable // Connect Redis cluster
```

## Usage Examples

### MySQL (GORM)

```go
package main

import (
	"github.com/zzhuang94/go-kit/db"
	"gorm.io/gorm"
)

func main() {
	cfg := &db.MysqlCfg{
		DSN:     "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdle: 10,
		MaxOpen: 100,
		Log:     true,
	}
	
	db, err := cfg.ConnGorm()
	if err != nil {
		panic(err)
	}
	// Use db...
}
```

### MySQL (XORM)

```go
package main

import (
	"github.com/zzhuang94/go-kit/db"
	"xorm.io/xorm"
)

func main() {
	cfg := &db.MysqlCfg{
		DSN:     "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4",
		MaxIdle: 10,
		MaxOpen: 100,
		Log:     true,
	}
	
	engine, err := cfg.ConnXorm()
	if err != nil {
		panic(err)
	}
	// Use engine...
}
```

### Redis

```go
package main

import (
	"github.com/zzhuang94/go-kit/db"
	"context"
)

func main() {
	cfg := &db.RedisCfg{
		Addr:   "localhost:6379",
		Passwd: "password",
	}
	
	client := db.ConnRedis(cfg)
	if client == nil {
		panic("Failed to connect to Redis")
	}
	
	// Use client...
	client.Set(context.Background(), "key", "value", 0)
}
```
