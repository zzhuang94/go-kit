package db

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

type RedisCfg struct {
	Addr   string `json:"addr"`
	Passwd string `json:"passwd"`
}

type RedisClusterCfg struct {
	Addrs  []string `json:"addrs"`
	Passwd string   `json:"passwd"`
}

// ConnRedis 连接单机 Redis / Connect standalone Redis
func ConnRedis(c *RedisCfg) redis.Cmdable {
	ret := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Passwd,
	})
	if err := ret.Ping(context.Background()).Err(); err != nil {
		log.Print(err.Error())
		return nil
	}
	return ret
}

// ConnRedisCluster 连接 Redis 集群 / Connect Redis cluster
func ConnRedisCluster(c *RedisClusterCfg) redis.Cmdable {
	ret := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    c.Addrs,
		Password: c.Passwd,
	})
	if err := ret.Ping(context.Background()).Err(); err != nil {
		log.Print(err.Error())
		return nil
	}
	return ret
}
