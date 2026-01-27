package lib

import (
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type lock struct {
	*limiter
}

func tryLock(s storage, key string, timeout time.Duration) (*lock, error) {
	l, err := tryCheckIn(s, key, 1, timeout)
	if err != nil {
		return nil, err
	}
	lock := new(lock)
	lock.limiter = l
	return lock, nil
}

func TryLockWithRedis(
	c redis.Cmdable,
	key string,
	timeout time.Duration,
) (*lock, error) {
	if c == nil {
		return nil, fmt.Errorf("redis client cannot be nil")
	}
	return tryLock(newRedisStorage(c), key, timeout)
}

func TryLockWithEtcd(
	client *clientv3.Client,
	key string,
	timeout time.Duration,
) (*lock, error) {
	if client == nil {
		return nil, fmt.Errorf("etcd client cannot be nil")
	}
	return tryLock(newEtcdStorage(client), key, timeout)
}

func TryLockLocal(key string, timeout time.Duration) (*lock, error) {
	return tryLock(newLocalStorage(), key, timeout)
}
