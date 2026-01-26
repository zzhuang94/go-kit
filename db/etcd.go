package db

import (
	"context"
	"log"
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
	timeout := 5 * time.Second

	config := clientv3.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: timeout,
	}

	if c.Username != "" && c.Password != "" {
		config.Username = c.Username
		config.Password = c.Password
	}

	client, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err = client.Status(ctx, c.Endpoints[0])
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		client.Close()
		return nil, err
	}

	return client, nil
}
