package db

import (
	redis "github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

func GetTestMysqlCfg() *MysqlCfg {
	return &MysqlCfg{
		DSN:     "root:Hzz123!@tcp(gin-vue.web.domain:3306)/hzz?charset=utf8&parseTime=true&loc=Local",
		MaxIdle: 10,
		MaxOpen: 100,
	}
}

func GetTestRedisCfg() *RedisCfg {
	return &RedisCfg{
		Addr:   "gin-vue.web.domain:6379",
		Passwd: "Hzz*Dev!",
	}
}

func GetTestEtcdCfg() *EtcdCfg {
	return &EtcdCfg{
		Endpoints: []string{"gin-vue.web.domain:2379"},
		Username:  "hzz",
		Password:  "Hzz321",
	}
}

func GetTestGorm() *gorm.DB {
	cfg := GetTestMysqlCfg()
	db, err := cfg.ConnGorm()
	if err != nil {
		panic(err)
	}
	return db
}

func GetTestXorm() *xorm.Engine {
	cfg := GetTestMysqlCfg()
	engine, err := cfg.ConnXorm()
	if err != nil {
		panic(err)
	}
	return engine
}

func GetTestRedis() redis.Cmdable {
	cfg := GetTestRedisCfg()
	return cfg.ConnRedis()
}

func GetTestEtcd() *clientv3.Client {
	cfg := GetTestEtcdCfg()
	client, err := cfg.ConnEtcd()
	if err != nil {
		panic(err)
	}
	return client
}
