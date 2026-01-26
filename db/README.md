# Database Utilities Package

提供数据库连接工具函数。

## MySQL 连接

### GORM 连接

```go
type MysqlCfg struct {
	DSN     string `json:"dsn"`      // 数据源名称
	MaxIdle int    `json:"max_idle"` // 最大空闲连接数
	MaxOpen int    `json:"max_open"` // 最大打开连接数
	Log     bool   `json:"log"`      // 是否开启日志
}

func (c *MysqlCfg) ConnGorm() (*gorm.DB, error) // 使用 GORM 连接 MySQL
```

### XORM 连接

```go
func (c *MysqlCfg) ConnXorm() (*xorm.Engine, error) // 使用 XORM 连接 MySQL
```

## Redis 连接

### 单机 Redis

```go
type RedisCfg struct {
	Addr   string `json:"addr"`   // Redis 地址
	Passwd string `json:"passwd"` // Redis 密码
}

func (c *RedisCfg) ConnRedis() redis.Cmdable // 连接单机 Redis
```

### Redis 集群

```go
type RedisClusterCfg struct {
	Addrs  []string `json:"addrs"`  // Redis 集群地址列表
	Passwd string   `json:"passwd"`  // Redis 密码
}

func (c *RedisClusterCfg) ConnRedisCluster() redis.Cmdable // 连接 Redis 集群
```

## etcd 连接

```go
type EtcdCfg struct {
	Endpoints   []string      `json:"endpoints"`    // etcd 端点地址列表
	Username    string        `json:"username"`     // 用户名（可选）
	Password    string        `json:"password"`     // 密码（可选）
}

func (c *EtcdCfg) ConnEtcd() (*clientv3.Client, error) // 连接 etcd，会自动测试连接
```

## 使用示例

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
	// 使用 db...
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
	// 使用 engine...
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
	
	// 使用 client...
	client.Set(context.Background(), "key", "value", 0)
}
```

### Redis 集群

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
	
	// 使用 client...
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
		Username:    "",                    // 可选，如果 etcd 启用了认证
		Password:    "",                    // 可选，如果 etcd 启用了认证
	}
	
	client, err := cfg.ConnEtcd()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	
	// 使用 client 进行 etcd 操作
	// 例如：使用 clientv3 的 API
	ctx := context.Background()
	_, err = client.Put(ctx, "key", "value")
	if err != nil {
		panic(err)
	}
	
	resp, err := client.Get(ctx, "key")
	if err != nil {
		panic(err)
	}
	
	// 使用 resp...
}
```
