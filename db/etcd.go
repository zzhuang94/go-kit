package db

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdCfg struct {
	Endpoints []string `json:"endpoints"` // etcd 端点地址列表
	Username  string   `json:"username"`  // 用户名（可选）
	Password  string   `json:"password"`  // 密码（可选）
}

// ConnEtcd 连接 etcd / Connect to etcd
func (c *EtcdCfg) ConnEtcd() (*clientv3.Client, error) {
	cfg := clientv3.Config{
		Endpoints:   c.Endpoints,
		Username:    c.Username,
		Password:    c.Password,
		DialTimeout: time.Second * 3,
	}
	return clientv3.New(cfg)
}
