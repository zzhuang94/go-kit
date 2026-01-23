package lib

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const (
	REDIS_VAL     int           = 1
	TTL           time.Duration = time.Second
	TRY_INTERVAL  time.Duration = 50 * time.Millisecond
	HOLD_INTERVAL time.Duration = 500 * time.Millisecond
)

var (
	mutex           sync.Mutex
	localLimiterMap map[string]int
)

type limiter struct {
	key          string
	releaseFlag  int32
	isLocal      bool
	redisCmdable redis.Cmdable
	releaseCtx   context.Context
	releaseFunc  context.CancelFunc
}

func TryCheckIn(c redis.Cmdable, key string, limit int, timeout time.Duration) (*limiter, error) {
	maxLimit := int(^uint(0) >> 1)
	if limit > maxLimit {
		return nil, fmt.Errorf("限制数不得大于[%d]", maxLimit)
	}

	timeLine := time.Now().Add(timeout)
	for {
		var l *limiter
		var err error
		if c != nil {
			l, err = tryCheckInRedis(c, key, limit)
		} else {
			l = tryCheckInLocal(key, limit)
		}
		if err != nil {
			return nil, err
		}
		if l != nil {
			return l, nil
		}
		if time.Now().After(timeLine) {
			return nil, fmt.Errorf("timeout")
		}
		time.Sleep(TRY_INTERVAL)
	}
}

func tryCheckInRedis(c redis.Cmdable, key string, limit int) (*limiter, error) {
	if err := c.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	res, err := c.Incr(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	if int(res) <= limit {
		l := new(limiter)
		l.key = key
		l.redisCmdable = c
		l.holding()
		return l, nil
	}
	c.Decr(context.Background(), key)
	return nil, nil
}

func tryCheckInLocal(key string, limit int) *limiter {
	mutex.Lock()
	defer mutex.Unlock()
	if localLimiterMap == nil {
		localLimiterMap = make(map[string]int)
	}
	res := localLimiterMap[key] + 1
	if res <= limit {
		localLimiterMap[key] = res
		l := new(limiter)
		l.key = key
		l.isLocal = true
		return l
	}
	return nil
}

func (l *limiter) Release() {
	if !atomic.CompareAndSwapInt32(&l.releaseFlag, 0, 1) {
		// 禁止重复release
		return
	}
	if l.isLocal {
		l.decrLocal()
	} else {
		l.releaseFunc()
		l.decrRedis()
	}
}

func (l *limiter) holding() {
	l.resetTTL()
	l.releaseCtx, l.releaseFunc = context.WithCancel(context.Background())
	ticker := time.NewTicker(HOLD_INTERVAL)
	isReleased := func() bool {
		select {
		case <-l.releaseCtx.Done():
			return true
		case <-ticker.C:
			return false
		}
	}
	go func() {
		for {
			if isReleased() {
				return
			}
			l.resetTTL()
		}
	}()
}

func (l *limiter) resetTTL() {
	l.redisCmdable.Expire(context.Background(), l.key, TTL)
}

func (l *limiter) decrRedis() {
	l.redisCmdable.Decr(context.Background(), l.key)
}

func (l *limiter) decrLocal() {
	mutex.Lock()
	defer mutex.Unlock()
	localLimiterMap[l.key] -= 1
}
