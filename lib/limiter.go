package lib

import (
	"context"
	"fmt"
	"strconv"
	"sync"
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

type storage interface {
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, ttl time.Duration) error
	Ping(ctx context.Context) error
}

type redisStorage struct {
	client redis.Cmdable
}

func (r *redisStorage) Incr(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

func (r *redisStorage) Decr(ctx context.Context, key string) error {
	return r.client.Decr(ctx, key).Err()
}

func (r *redisStorage) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return r.client.Expire(ctx, key, ttl).Err()
}

func (r *redisStorage) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func newRedisStorage(c redis.Cmdable) storage {
	return &redisStorage{client: c}
}

type localStorage struct {
	mutex sync.Mutex
	data  map[string]int64
}

func (l *localStorage) Incr(ctx context.Context, key string) (int64, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.data == nil {
		l.data = make(map[string]int64)
	}
	l.data[key]++
	return l.data[key], nil
}

func (l *localStorage) Decr(ctx context.Context, key string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.data == nil {
		return nil
	}
	if val, ok := l.data[key]; ok && val > 0 {
		l.data[key]--
	}
	return nil
}

func (l *localStorage) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return nil
}

func (l *localStorage) Ping(ctx context.Context) error {
	return nil
}

func newLocalStorage() storage {
	return &localStorage{}
}

type etcdStorage struct {
	client   *clientv3.Client
	leaseMap sync.Map
}

func (e *etcdStorage) Incr(ctx context.Context, key string) (int64, error) {
	for {
		resp, err := e.client.Get(ctx, key)
		if err != nil {
			return 0, err
		}
		var val int64
		if len(resp.Kvs) > 0 {
			val, err = strconv.ParseInt(string(resp.Kvs[0].Value), 10, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse value: %w", err)
			}
		}
		nval := val + 1
		valStr := strconv.FormatInt(nval, 10)
		var lid clientv3.LeaseID
		if lidVal, ok := e.leaseMap.Load(key); ok {
			lid = lidVal.(clientv3.LeaseID)
		} else {
			lease, err := e.client.Grant(ctx, int64(TTL.Seconds()))
			if err != nil {
				return 0, err
			}
			lid = lease.ID
			e.leaseMap.Store(key, lid)
		}
		txn := e.client.Txn(ctx)
		if len(resp.Kvs) > 0 {
			txn = txn.If(
				clientv3.Compare(clientv3.Version(key), "=", resp.Kvs[0].Version),
			)
		} else {
			txn = txn.If(clientv3.Compare(clientv3.Version(key), "=", 0))
		}
		txn = txn.Then(
			clientv3.OpPut(key, valStr, clientv3.WithLease(lid)),
		)
		txn = txn.Else(clientv3.OpGet(key))
		tr, err := txn.Commit()
		if err != nil {
			return 0, err
		}
		if tr.Succeeded {
			return nval, nil
		}
	}
}

func (e *etcdStorage) Decr(ctx context.Context, key string) error {
	for {
		resp, err := e.client.Get(ctx, key)
		if err != nil {
			return err
		}
		if len(resp.Kvs) == 0 {
			e.leaseMap.Delete(key)
			return nil
		}
		val, err := strconv.ParseInt(string(resp.Kvs[0].Value), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse value: %w", err)
		}
		if val <= 0 {
			return nil
		}
		nval := val - 1
		valStr := strconv.FormatInt(nval, 10)
		var lid clientv3.LeaseID
		if lidVal, ok := e.leaseMap.Load(key); ok {
			lid = lidVal.(clientv3.LeaseID)
		} else {
			lease, err := e.client.Grant(ctx, int64(TTL.Seconds()))
			if err != nil {
				return err
			}
			lid = lease.ID
			e.leaseMap.Store(key, lid)
		}
		txn := e.client.Txn(ctx)
		txn = txn.If(
			clientv3.Compare(clientv3.Version(key), "=", resp.Kvs[0].Version),
		)
		txn = txn.Then(
			clientv3.OpPut(key, valStr, clientv3.WithLease(lid)),
		)
		txn = txn.Else(clientv3.OpGet(key))
		tr, err := txn.Commit()
		if err != nil {
			return err
		}
		if tr.Succeeded {
			return nil
		}
	}
}

func (e *etcdStorage) Expire(ctx context.Context, key string, ttl time.Duration) error {
	lidVal, ok := e.leaseMap.Load(key)
	if !ok {
		resp, err := e.client.Get(ctx, key)
		if err != nil {
			return err
		}
		if len(resp.Kvs) == 0 {
			return nil
		}
		lease, err := e.client.Grant(ctx, int64(ttl.Seconds()))
		if err != nil {
			return err
		}
		e.leaseMap.Store(key, lease.ID)
		txn := e.client.Txn(ctx)
		txn = txn.If(
			clientv3.Compare(clientv3.Version(key), "=", resp.Kvs[0].Version),
		)
		txn = txn.Then(
			clientv3.OpPut(
				key,
				string(resp.Kvs[0].Value),
				clientv3.WithLease(lease.ID),
			),
		)
		tr, err := txn.Commit()
		if err != nil {
			return err
		}
		if !tr.Succeeded {
			e.leaseMap.Delete(key)
			return nil
		}
		lidVal = lease.ID
	}
	lid := lidVal.(clientv3.LeaseID)
	_, err := e.client.KeepAliveOnce(ctx, lid)
	return err
}

func (e *etcdStorage) Ping(ctx context.Context) error {
	if len(e.client.Endpoints()) == 0 {
		return fmt.Errorf("no etcd endpoints configured")
	}
	_, err := e.client.Status(ctx, e.client.Endpoints()[0])
	return err
}

func newEtcdStorage(client *clientv3.Client) storage {
	return &etcdStorage{client: client}
}

type limiter struct {
	key      string
	released int32
	storage  storage
	ctx      context.Context
	cancel   context.CancelFunc
}

func tryCheckIn(s storage, key string, limit int, timeout time.Duration) (*limiter, error) {
	maxLimit := int(^uint(0) >> 1)
	if limit > maxLimit {
		return nil, fmt.Errorf("限制数不得大于[%d]", maxLimit)
	}
	deadline := time.Now().Add(timeout)
	for {
		l, err := tryOnce(s, key, limit)
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

func tryOnce(s storage, key string, limit int) (*limiter, error) {
	ctx := context.Background()
	if err := s.Ping(ctx); err != nil {
		return nil, err
	}
	res, err := s.Incr(ctx, key)
	if err != nil {
		return nil, err
	}
	if int(res) <= limit {
		l := new(limiter)
		l.key = key
		l.storage = s
		l.holding()
		return l, nil
	}
	s.Decr(ctx, key)
	return nil, nil
}

func TryCheckInWithRedis(c redis.Cmdable, k string, l int, t time.Duration) (*limiter, error) {
	if c == nil {
		return nil, fmt.Errorf("redis client cannot be nil")
	}
	return tryCheckIn(newRedisStorage(c), k, l, t)
}

func TryCheckInWithEtcd(c *clientv3.Client, k string, l int, t time.Duration) (*limiter, error) {
	if c == nil {
		return nil, fmt.Errorf("etcd client cannot be nil")
	}
	return tryCheckIn(newEtcdStorage(c), k, l, t)
}

func TryCheckInLocal(k string, l int, t time.Duration) (*limiter, error) {
	return tryCheckIn(newLocalStorage(), k, l, t)
}

func (l *limiter) Release() {
	if !atomic.CompareAndSwapInt32(&l.released, 0, 1) {
		return
	}
	if l.cancel != nil {
		l.cancel()
	}
	l.decr()
}

func (l *limiter) holding() {
	l.resetTTL()
	l.ctx, l.cancel = context.WithCancel(context.Background())
	ticker := time.NewTicker(HOLD_INTERVAL)
	done := func() bool {
		select {
		case <-l.ctx.Done():
			return true
		case <-ticker.C:
			return false
		}
	}
	go func() {
		defer ticker.Stop()
		for {
			if done() {
				return
			}
			l.resetTTL()
		}
	}()
}

func (l *limiter) resetTTL() {
	if l.storage != nil {
		l.storage.Expire(context.Background(), l.key, TTL)
	}
}

func (l *limiter) decr() {
	if l.storage != nil {
		l.storage.Decr(context.Background(), l.key)
	}
}
