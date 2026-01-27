package lib

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type KeyStor interface {
	Incr(ctx context.Context) (int64, error)
	Decr(ctx context.Context) error
	Expire(ctx context.Context, ttl time.Duration) error
	Ping(ctx context.Context) error
}

var (
	localMutex sync.Mutex
	localMap   map[string]int64
)

type KeyStorLocal struct {
	key string
}

func NewKeyStorLocal(key string) KeyStor {
	localMutex.Lock()
	defer localMutex.Unlock()
	if localMap == nil {
		localMap = make(map[string]int64)
	}
	if _, ok := localMap[key]; !ok {
		localMap[key] = 0
	}
	return &KeyStorLocal{key: key}
}

func (l *KeyStorLocal) Incr(ctx context.Context) (int64, error) {
	localMutex.Lock()
	defer localMutex.Unlock()
	localMap[l.key]++
	return localMap[l.key], nil
}

func (l *KeyStorLocal) Decr(ctx context.Context) error {
	localMutex.Lock()
	defer localMutex.Unlock()
	if localMap[l.key] > 0 {
		localMap[l.key]--
	}
	return nil
}

func (l *KeyStorLocal) Expire(ctx context.Context, ttl time.Duration) error {
	return nil
}

func (l *KeyStorLocal) Ping(ctx context.Context) error {
	return nil
}

type KeyStorRedis struct {
	client redis.Cmdable
	key    string
}

func NewKeyStorRedis(c redis.Cmdable, key string) KeyStor {
	return &KeyStorRedis{client: c, key: key}
}
func (r *KeyStorRedis) Incr(ctx context.Context) (int64, error) {
	return r.client.Incr(ctx, r.key).Result()
}

func (r *KeyStorRedis) Decr(ctx context.Context) error {
	return r.client.Decr(ctx, r.key).Err()
}

func (r *KeyStorRedis) Expire(ctx context.Context, ttl time.Duration) error {
	return r.client.Expire(ctx, r.key, ttl).Err()
}

func (r *KeyStorRedis) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

type KeyStorEtcd struct {
	client *clientv3.Client
	key    string
	lease  clientv3.LeaseID
}

func (e *KeyStorEtcd) Incr(ctx context.Context) (int64, error) {
	for {
		resp, err := e.client.Get(ctx, e.key)
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
		if e.lease == 0 {
			lease, err := e.client.Grant(ctx, int64(TTL.Seconds()))
			if err != nil {
				return 0, err
			}
			e.lease = lease.ID
		}
		lid := e.lease
		txn := e.client.Txn(ctx)
		if len(resp.Kvs) > 0 {
			txn = txn.If(
				clientv3.Compare(clientv3.Version(e.key), "=", resp.Kvs[0].Version),
			)
		} else {
			txn = txn.If(clientv3.Compare(clientv3.Version(e.key), "=", 0))
		}
		txn = txn.Then(
			clientv3.OpPut(e.key, valStr, clientv3.WithLease(lid)),
		)
		txn = txn.Else(clientv3.OpGet(e.key))
		tr, err := txn.Commit()
		if err != nil {
			return 0, err
		}
		if tr.Succeeded {
			return nval, nil
		}
	}
}

func (e *KeyStorEtcd) Decr(ctx context.Context) error {
	for {
		resp, err := e.client.Get(ctx, e.key)
		if err != nil {
			return err
		}
		if len(resp.Kvs) == 0 {
			e.lease = 0
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
		if e.lease == 0 {
			lease, err := e.client.Grant(ctx, int64(TTL.Seconds()))
			if err != nil {
				return err
			}
			e.lease = lease.ID
		}
		lid := e.lease
		txn := e.client.Txn(ctx)
		txn = txn.If(
			clientv3.Compare(clientv3.Version(e.key), "=", resp.Kvs[0].Version),
		)
		txn = txn.Then(
			clientv3.OpPut(e.key, valStr, clientv3.WithLease(lid)),
		)
		txn = txn.Else(clientv3.OpGet(e.key))
		tr, err := txn.Commit()
		if err != nil {
			return err
		}
		if tr.Succeeded {
			return nil
		}
	}
}

func (e *KeyStorEtcd) Expire(ctx context.Context, ttl time.Duration) error {
	if e.lease == 0 {
		resp, err := e.client.Get(ctx, e.key)
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
		e.lease = lease.ID
		txn := e.client.Txn(ctx)
		txn = txn.If(
			clientv3.Compare(clientv3.Version(e.key), "=", resp.Kvs[0].Version),
		)
		txn = txn.Then(
			clientv3.OpPut(
				e.key,
				string(resp.Kvs[0].Value),
				clientv3.WithLease(lease.ID),
			),
		)
		tr, err := txn.Commit()
		if err != nil {
			return err
		}
		if !tr.Succeeded {
			e.lease = 0
			return nil
		}
	}
	_, err := e.client.KeepAliveOnce(ctx, e.lease)
	return err
}

func (e *KeyStorEtcd) Ping(ctx context.Context) error {
	if len(e.client.Endpoints()) == 0 {
		return fmt.Errorf("no etcd endpoints configured")
	}
	_, err := e.client.Status(ctx, e.client.Endpoints()[0])
	return err
}

func NewKeyStorEtcd(client *clientv3.Client, key string) KeyStor {
	return &KeyStorEtcd{client: client, key: key}
}
