package lib

import (
	"time"

	redis "github.com/go-redis/redis/v8"
)

type lock struct {
	*limiter
}

func TryLock(c redis.Cmdable, key string, timeout time.Duration) (*lock, error) {
	limiter, err := TryCheckIn(c, key, 1, timeout)
	if err != nil {
		return nil, err
	}
	lock := new(lock)
	lock.limiter = limiter
	return lock, nil
}
