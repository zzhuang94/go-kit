package lib

import (
	"context"
	"log"
	"strings"
	"sync/atomic"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/zzhuang94/go-kit/str"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	eleTTL  = time.Second
	eleTICK = time.Millisecond * 500
)

type Election struct {
	eleStor
	isMaster *int32
	val      string
	uuid     string
	ctx      context.Context
	cancel   context.CancelFunc
}

type eleStor interface {
	Key() string
	Get(ctx context.Context) (string, error)
	Del(ctx context.Context) error
	SetNX(ctx context.Context, value string, ttl time.Duration) (bool, error)
	Expire(ctx context.Context, ttl time.Duration) error
}

func NewElectionWithRedis(cmdable redis.Cmdable, key, val string) *Election {
	if cmdable == nil {
		panic("redis client cannot be nil")
	}
	return newElection(NewRedisStor(cmdable, key), val)
}

func NewElectionWithEtcd(client *clientv3.Client, key, val string) *Election {
	if client == nil {
		panic("etcd client cannot be nil")
	}
	return newElection(NewEtcdStor(client, key), val)
}

func newElection(stor eleStor, val string) *Election {
	e := &Election{
		eleStor:  stor,
		isMaster: new(int32),
		val:      val,
		uuid:     str.Uuid(),
	}
	e.start()
	return e
}

func (e *Election) start() {
	e.ctx, e.cancel = context.WithCancel(context.Background())
	e.checkIsMaster()
	go e.polling()
}

func (e *Election) polling() {
	ticker := time.NewTicker(eleTICK)
	for {
		defer ticker.Stop()
		select {
		case <-e.ctx.Done():
			return
		case <-ticker.C:
			e.checkIsMaster()
		}
	}
}

func (e *Election) checkIsMaster() {
	ctx := context.Background()
	val := e.buildVal()
	if e.IsMaster() {
		currentVal, err := e.Get(ctx)
		if err != nil || val != currentVal {
			atomic.StoreInt32(e.isMaster, 0)
		}
	} else {
		success, err := e.SetNX(ctx, val, eleTTL)
		if err == nil && success {
			log.Printf("%s GET MASTER", e.val)
			atomic.StoreInt32(e.isMaster, 1)
		}
	}
	if e.IsMaster() {
		e.Expire(ctx, eleTTL)
	}
}

func (e *Election) buildVal() string {
	return e.uuid + "|" + e.val
}

func (e *Election) getVal() string {
	v, _ := e.Get(e.ctx)
	return v
}

func (e *Election) IsMaster() bool {
	return atomic.LoadInt32(e.isMaster) == 1
}

func (e *Election) Release() {
	if !e.IsMaster() {
		return
	}
	atomic.StoreInt32(e.isMaster, 0)
	e.Del(e.ctx)
}

func (e *Election) GetMaster() string {
	v := e.getVal()
	if v == "" {
		return ""
	}
	ss := strings.Split(v, "|")
	if len(ss) < 2 {
		return ""
	}
	return v[33:]
}

func (e *Election) Stop() {
	e.Release()
	e.cancel()
}
