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

var (
	localMutex sync.Mutex
	localMap   map[string]int64
)

type LocalStor struct {
	key string
}

func NewLocalStor(key string) *LocalStor {
	localMutex.Lock()
	defer localMutex.Unlock()
	if localMap == nil {
		localMap = make(map[string]int64)
	}
	if _, ok := localMap[key]; !ok {
		localMap[key] = 0
	}
	return &LocalStor{key: key}
}

func (l *LocalStor) Key() string {
	return l.key
}

func (l *LocalStor) Incr(ctx context.Context) (int64, error) {
	localMutex.Lock()
	defer localMutex.Unlock()
	localMap[l.key]++
	return localMap[l.key], nil
}

func (l *LocalStor) Decr(ctx context.Context) error {
	localMutex.Lock()
	defer localMutex.Unlock()
	if localMap[l.key] > 0 {
		localMap[l.key]--
	}
	return nil
}

func (l *LocalStor) Expire(ctx context.Context, ttl time.Duration) error {
	return nil
}

func (l *LocalStor) Ping(ctx context.Context) error {
	return nil
}

type RedisStor struct {
	client redis.Cmdable
	key    string
}

func NewRedisStor(c redis.Cmdable, key string) *RedisStor {
	return &RedisStor{client: c, key: key}
}

func (r *RedisStor) Key() string {
	return r.key
}

func (r *RedisStor) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *RedisStor) Get(ctx context.Context) (string, error) {
	return r.client.Get(ctx, r.key).Result()
}

func (r *RedisStor) Del(ctx context.Context) error {
	return r.client.Del(ctx, r.key).Err()
}

func (r *RedisStor) Incr(ctx context.Context) (int64, error) {
	return r.client.Incr(ctx, r.key).Result()
}

func (r *RedisStor) Decr(ctx context.Context) error {
	return r.client.Decr(ctx, r.key).Err()
}

func (r *RedisStor) Expire(ctx context.Context, ttl time.Duration) error {
	return r.client.Expire(ctx, r.key, ttl).Err()
}

func (r *RedisStor) SetNX(ctx context.Context, value string, ttl time.Duration) (bool, error) {
	return r.client.SetNX(ctx, r.key, value, ttl).Result()
}

type EtcdStor struct {
	client *clientv3.Client
	key    string
	lease  clientv3.LeaseID
}

func NewEtcdStor(client *clientv3.Client, key string) *EtcdStor {
	return &EtcdStor{client: client, key: key}
}

func (e *EtcdStor) Key() string {
	return e.key
}

func (e *EtcdStor) Incr(ctx context.Context) (int64, error) {
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

func (e *EtcdStor) Decr(ctx context.Context) error {
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

func (e *EtcdStor) Expire(ctx context.Context, ttl time.Duration) error {
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

func (e *EtcdStor) Ping(ctx context.Context) error {
	_, err := e.client.Get(ctx, e.key)
	return err
}

func (e *EtcdStor) SetNX(ctx context.Context, value string, ttl time.Duration) (bool, error) {
	// Create lease for TTL first
	lease, err := e.client.Grant(ctx, int64(ttl.Seconds()))
	if err != nil {
		return false, err
	}

	// Use transaction to atomically set if not exists
	txn := e.client.Txn(ctx)
	txn = txn.If(clientv3.Compare(clientv3.Version(e.key), "=", 0))
	txn = txn.Then(clientv3.OpPut(e.key, value, clientv3.WithLease(lease.ID)))
	txn = txn.Else(clientv3.OpGet(e.key))

	tr, err := txn.Commit()
	if err != nil {
		e.client.Revoke(ctx, lease.ID)
		return false, err
	}

	if tr.Succeeded {
		// Successfully set the key, store the lease ID for future Expire calls
		e.lease = lease.ID
		return true, nil
	}

	// Transaction failed, key already exists
	// Revoke the lease we created since we didn't use it
	e.client.Revoke(ctx, lease.ID)
	return false, nil
}

func (e *EtcdStor) Get(ctx context.Context) (string, error) {
	resp, err := e.client.Get(ctx, e.key)
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}

func (e *EtcdStor) Del(ctx context.Context) error {
	if e.lease != 0 {
		// Revoke lease first
		_, err := e.client.Revoke(ctx, e.lease)
		if err != nil {
			return err
		}
		e.lease = 0
	}
	_, err := e.client.Delete(ctx, e.key)
	return err
}
