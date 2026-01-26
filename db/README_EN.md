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

func (c *RedisCfg) ConnRedis() redis.Cmdable // Connect standalone Redis
```

### Redis Cluster

```go
type RedisClusterCfg struct {
	Addrs  []string `json:"addrs"`  // Redis cluster addresses
	Passwd string   `json:"passwd"`  // Redis password
}

func (c *RedisClusterCfg) ConnRedisCluster() redis.Cmdable // Connect Redis cluster
```

## etcd Connection

```go
type EtcdCfg struct {
	Endpoints   []string      `json:"endpoints"`    // etcd endpoint addresses
	Username    string        `json:"username"`     // Username (optional)
	Password    string        `json:"password"`     // Password (optional)seconds
}

func (c *EtcdCfg) ConnEtcd() (*clientv3.Client, error) // Connect to etcd, automatically tests the connection
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
	
	client := cfg.ConnRedis()
	if client == nil {
		panic("Failed to connect to Redis")
	}
	
	// Use client...
	client.Set(context.Background(), "key", "value", 0)
}
```

### Redis Cluster

```go
package main

import (
	"github.com/zzhuang94/go-kit/db"
	"context"
)

func main() {
	cfg := &db.RedisClusterCfg{
		Addrs:  []string{"localhost:7000", "localhost:7001", "localhost:7002"},
		Passwd: "password",
	}
	
	client := cfg.ConnRedisCluster()
	if client == nil {
		panic("Failed to connect to Redis cluster")
	}
	
	// Use client...
	client.Set(context.Background(), "key", "value", 0)
}
```

### etcd

```go
package main

import (
	"context"
	"github.com/zzhuang94/go-kit/db"
	"time"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cfg := &db.EtcdCfg{
		Endpoints:   []string{"localhost:2379"},
		Username:    "",                    // Optional, if etcd has authentication enabled
		Password:    "",                    // Optional, if etcd has authentication enabled
	}
	
	client, err := cfg.ConnEtcd()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	
	// Use client for etcd operations
	// For example: using clientv3 API
	ctx := context.Background()
	_, err = client.Put(ctx, "key", "value")
	if err != nil {
		panic(err)
	}
	
	resp, err := client.Get(ctx, "key")
	if err != nil {
		panic(err)
	}
	
	// Use resp...
}
```
