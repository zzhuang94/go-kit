package lib

import (
	"context"
	"log"
	"strings"
	"sync/atomic"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/zzhuang94/go-kit/str"
)

const (
	ELE_TTL  = time.Second
	ELE_TICK = time.Millisecond * 500
)

type Election struct {
	isMaster *int32
	key      string
	val      string
	uuid     string
	ctx      context.Context
	cancel   context.CancelFunc
	cmdable  redis.Cmdable
}

func NewElection(key, val string, cmdable redis.Cmdable) *Election {
	e := &Election{
		isMaster: new(int32),
		key:      key,
		val:      val,
		uuid:     str.Uuid(),
		cmdable:  cmdable,
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
	ticker := time.NewTicker(ELE_TICK)
	for {
		select {
		case <-e.ctx.Done():
			atomic.StoreInt32(e.isMaster, 0)
			ticker.Stop()
			return
		case <-ticker.C:
			e.checkIsMaster()
		}
	}

}

func (e *Election) checkIsMaster() {
	redisVal := e.buildRedisVal()
	if e.IsMaster() {
		if redisVal != e.getRedisVal() {
			atomic.StoreInt32(e.isMaster, 0)
		}
	} else {
		if e.cmdable.SetNX(e.ctx, e.key, redisVal, ELE_TTL).Val() {
			log.Printf("%s GET MASTER", e.val)
			atomic.StoreInt32(e.isMaster, 1)
		}
	}
	if e.IsMaster() {
		e.cmdable.Expire(e.ctx, e.key, ELE_TTL)
	}
}

func (e *Election) buildRedisVal() string {
	return e.uuid + "|" + e.val
}

func (e *Election) getRedisVal() string {
	v, _ := e.cmdable.Get(e.ctx, e.key).Result()
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
	e.cmdable.Del(e.ctx, e.key)
}

func (e *Election) GetMaster() string {
	v := e.getRedisVal()
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
