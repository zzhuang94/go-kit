package lib

import (
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Lock struct {
	*Limiter
}

func tryLock(s limStor, timeout time.Duration) (*Lock, error) {
	l, err := tryCheckIn(s, 1, timeout)
	if err != nil {
		return nil, err
	}
	lock := new(Lock)
	lock.Limiter = l
	return lock, nil
}

func TryLockWithRedis(c redis.Cmdable, key string, timeout time.Duration) (*Lock, error) {
	if c == nil {
		return nil, fmt.Errorf("redis client cannot be nil")
	}
	return tryLock(NewRedisStor(c, key), timeout)
}

func TryLockWithEtcd(c *clientv3.Client, key string, timeout time.Duration) (*Lock, error) {
	if c == nil {
		return nil, fmt.Errorf("etcd client cannot be nil")
	}
	return tryLock(NewEtcdStor(c, key), timeout)
}

func TryLockLocal(key string, timeout time.Duration) (*Lock, error) {
	return tryLock(NewLocalStor(key), timeout)
}
