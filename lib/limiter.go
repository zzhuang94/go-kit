package lib

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	redis "github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	TTL           time.Duration = time.Second
	TRY_INTERVAL  time.Duration = 50 * time.Millisecond
	HOLD_INTERVAL time.Duration = 500 * time.Millisecond
)

type limiter struct {
	KeyStor
	released int32
	ctx      context.Context
	cancel   context.CancelFunc
}

func tryCheckIn(s KeyStor, limit int, timeout time.Duration) (*limiter, error) {
	maxLimit := int(^uint(0) >> 1)
	if limit > maxLimit {
		return nil, fmt.Errorf("限制数不得大于[%d]", maxLimit)
	}
	deadline := time.Now().Add(timeout)
	for {
		l, err := tryOnce(s, limit)
		if err != nil {
			return nil, err
		}
		if l != nil {
			return l, nil
		}
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("timeout")
		}
		time.Sleep(TRY_INTERVAL)
	}
}

func tryOnce(s KeyStor, limit int) (*limiter, error) {
	ctx := context.Background()
	if err := s.Ping(ctx); err != nil {
		return nil, err
	}
	res, err := s.Incr(ctx)
	if err != nil {
		return nil, err
	}
	if int(res) <= limit {
		l := new(limiter)
		l.KeyStor = s
		l.ctx, l.cancel = context.WithCancel(context.Background())
		l.holding()
		return l, nil
	}
	s.Decr(ctx)
	return nil, nil
}

func TryCheckInWithRedis(c redis.Cmdable, k string, l int, t time.Duration) (*limiter, error) {
	if c == nil {
		return nil, fmt.Errorf("redis client cannot be nil")
	}
	return tryCheckIn(NewKeyStorRedis(c, k), l, t)
}

func TryCheckInWithEtcd(c *clientv3.Client, k string, l int, t time.Duration) (*limiter, error) {
	if c == nil {
		return nil, fmt.Errorf("etcd client cannot be nil")
	}
	return tryCheckIn(NewKeyStorEtcd(c, k), l, t)
}

func TryCheckInLocal(k string, l int, t time.Duration) (*limiter, error) {
	return tryCheckIn(NewKeyStorLocal(k), l, t)
}

func (l *limiter) Release() {
	if !atomic.CompareAndSwapInt32(&l.released, 0, 1) {
		return
	}
	if l.cancel != nil {
		l.cancel()
	}
	l.KeyStor.Decr(context.Background())
}

func (l *limiter) holding() {
	l.expire()

	ticker := time.NewTicker(HOLD_INTERVAL)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-l.ctx.Done():
				return
			case <-ticker.C:
				l.expire()
			}
		}
	}()
}

func (l *limiter) expire() {
	l.KeyStor.Expire(context.Background(), TTL)
}
