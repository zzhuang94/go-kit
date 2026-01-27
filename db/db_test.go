package db

import (
	"context"
	"testing"
	"time"

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

// GoKitTest 测试表结构
type GoKitTest struct {
	ID        int       `gorm:"primaryKey;autoIncrement" xorm:"pk autoincr 'id'"`
	Name      string    `gorm:"type:varchar(100);not null" xorm:"varchar(100) not null 'name'"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex" xorm:"varchar(100) unique 'email'"`
	Age       int       `gorm:"type:int;default:0" xorm:"int default 0 'age'"`
	CreatedAt time.Time `gorm:"autoCreateTime" xorm:"created 'created_at'"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" xorm:"updated 'updated_at'"`
}

func (GoKitTest) TableName() string {
	return "go_kit_test"
}

// setupTestTable 设置测试表（删除已存在的表，然后创建新表）
func setupTestTable(t *testing.T) {
	gormDB := GetTestGorm()
	xormDB := GetTestXorm()

	// 删除已存在的表（如果存在）
	gormDB.Exec("DROP TABLE IF EXISTS go_kit_test")
	xormDB.Exec("DROP TABLE IF EXISTS go_kit_test")

	// 使用 GORM 创建表
	err := gormDB.AutoMigrate(&GoKitTest{})
	if err != nil {
		t.Fatalf("Failed to create table with GORM: %v", err)
	}

	t.Log("Test table 'go_kit_test' created successfully")
}

// TestMysqlGorm 测试 GORM 连接和基本操作
func TestMysqlGorm(t *testing.T) {
	// 设置测试表
	setupTestTable(t)

	db := GetTestGorm()
	if db == nil {
		t.Fatal("Failed to get GORM connection")
	}

	// 测试插入
	testUser := GoKitTest{
		Name:  "Test User",
		Email: "test@example.com",
		Age:   25,
	}
	result := db.Create(&testUser)
	if result.Error != nil {
		t.Fatalf("Failed to insert record: %v", result.Error)
	}
	if testUser.ID == 0 {
		t.Fatal("ID should be auto-generated")
	}
	t.Logf("Inserted record with ID: %d", testUser.ID)

	// 测试查询
	var foundUser GoKitTest
	result = db.First(&foundUser, testUser.ID)
	if result.Error != nil {
		t.Fatalf("Failed to query record: %v", result.Error)
	}
	if foundUser.Name != testUser.Name {
		t.Errorf("Expected name %s, got %s", testUser.Name, foundUser.Name)
	}
	if foundUser.Email != testUser.Email {
		t.Errorf("Expected email %s, got %s", testUser.Email, foundUser.Email)
	}
	if foundUser.Age != testUser.Age {
		t.Errorf("Expected age %d, got %d", testUser.Age, foundUser.Age)
	}
	t.Log("Query test passed")

	// 测试更新
	testUser.Age = 30
	testUser.Name = "Updated User"
	result = db.Save(&testUser)
	if result.Error != nil {
		t.Fatalf("Failed to update record: %v", result.Error)
	}

	var updatedUser GoKitTest
	db.First(&updatedUser, testUser.ID)
	if updatedUser.Age != 30 {
		t.Errorf("Expected age 30, got %d", updatedUser.Age)
	}
	if updatedUser.Name != "Updated User" {
		t.Errorf("Expected name 'Updated User', got %s", updatedUser.Name)
	}
	t.Log("Update test passed")

	// 测试计数
	var count int64
	db.Model(&GoKitTest{}).Count(&count)
	if count < 1 {
		t.Errorf("Expected at least 1 record, got %d", count)
	}
	t.Logf("Count test passed: %d records found", count)

	// 测试删除（但表保留，按用户要求）
	// result = db.Delete(&testUser)
	// if result.Error != nil {
	// 	t.Fatalf("Failed to delete record: %v", result.Error)
	// }
	t.Log("MySQL GORM tests completed successfully")
}

// TestMysqlXorm 测试 XORM 连接和基本操作
func TestMysqlXorm(t *testing.T) {
	// 设置测试表
	setupTestTable(t)

	engine := GetTestXorm()
	if engine == nil {
		t.Fatal("Failed to get XORM connection")
	}

	// 测试插入
	testUser := GoKitTest{
		Name:  "XORM Test User",
		Email: "xorm@example.com",
		Age:   28,
	}
	affected, err := engine.Insert(&testUser)
	if err != nil {
		t.Fatalf("Failed to insert record: %v", err)
	}
	if affected != 1 {
		t.Errorf("Expected 1 affected row, got %d", affected)
	}
	if testUser.ID == 0 {
		t.Fatal("ID should be auto-generated")
	}
	t.Logf("Inserted record with ID: %d", testUser.ID)

	// 测试查询
	var foundUser GoKitTest
	has, err := engine.ID(testUser.ID).Get(&foundUser)
	if err != nil {
		t.Fatalf("Failed to query record: %v", err)
	}
	if !has {
		t.Fatal("Record should exist")
	}
	if foundUser.Name != testUser.Name {
		t.Errorf("Expected name %s, got %s", testUser.Name, foundUser.Name)
	}
	if foundUser.Email != testUser.Email {
		t.Errorf("Expected email %s, got %s", testUser.Email, foundUser.Email)
	}
	if foundUser.Age != testUser.Age {
		t.Errorf("Expected age %d, got %d", testUser.Age, foundUser.Age)
	}
	t.Log("Query test passed")

	// 测试更新
	testUser.Age = 35
	testUser.Name = "XORM Updated User"
	affected, err = engine.ID(testUser.ID).Update(&testUser)
	if err != nil {
		t.Fatalf("Failed to update record: %v", err)
	}
	if affected != 1 {
		t.Errorf("Expected 1 affected row, got %d", affected)
	}

	var updatedUser GoKitTest
	engine.ID(testUser.ID).Get(&updatedUser)
	if updatedUser.Age != 35 {
		t.Errorf("Expected age 35, got %d", updatedUser.Age)
	}
	if updatedUser.Name != "XORM Updated User" {
		t.Errorf("Expected name 'XORM Updated User', got %s", updatedUser.Name)
	}
	t.Log("Update test passed")

	// 测试计数
	count, err := engine.Count(&GoKitTest{})
	if err != nil {
		t.Fatalf("Failed to count records: %v", err)
	}
	if count < 1 {
		t.Errorf("Expected at least 1 record, got %d", count)
	}
	t.Logf("Count test passed: %d records found", count)

	// 测试查询所有记录
	var users []GoKitTest
	err = engine.Find(&users)
	if err != nil {
		t.Fatalf("Failed to find all records: %v", err)
	}
	if len(users) < 1 {
		t.Errorf("Expected at least 1 record, got %d", len(users))
	}
	t.Logf("Find all test passed: %d records found", len(users))

	// 测试删除（但表保留，按用户要求）
	// affected, err = engine.ID(testUser.ID).Delete(&GoKitTest{})
	// if err != nil {
	// 	t.Fatalf("Failed to delete record: %v", err)
	// }
	t.Log("MySQL XORM tests completed successfully")
}

// TestMysqlConnection 测试 MySQL 连接
func TestMysqlConnection(t *testing.T) {
	cfg := GetTestMysqlCfg()

	// 测试 GORM 连接
	gormDB, err := cfg.ConnGorm()
	if err != nil {
		t.Fatalf("Failed to connect with GORM: %v", err)
	}
	sqlDB, _ := gormDB.DB()
	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Failed to ping database with GORM: %v", err)
	}
	t.Log("GORM connection test passed")

	// 测试 XORM 连接
	xormDB, err := cfg.ConnXorm()
	if err != nil {
		t.Fatalf("Failed to connect with XORM: %v", err)
	}
	if err := xormDB.Ping(); err != nil {
		t.Fatalf("Failed to ping database with XORM: %v", err)
	}
	t.Log("XORM connection test passed")
}

// TestRedisConnection 测试 Redis 连接
func TestRedisConnection(t *testing.T) {
	cfg := GetTestRedisCfg()
	client := cfg.ConnRedis()
	if client == nil {
		t.Skip("Redis is not available, skipping Redis tests")
		return
	}

	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		t.Fatalf("Failed to ping Redis: %v", err)
	}
	t.Log("Redis connection test passed")
}

// TestRedisBasicOperations 测试 Redis 基本操作
func TestRedisBasicOperations(t *testing.T) {
	client := GetTestRedis()
	if client == nil {
		t.Skip("Redis is not available, skipping Redis tests")
		return
	}

	ctx := context.Background()
	testKey := "go_kit_test:basic"
	testValue := "test_value_123"

	// 清理测试数据
	client.Del(ctx, testKey)

	// 测试 SET
	err := client.Set(ctx, testKey, testValue, 0).Err()
	if err != nil {
		t.Fatalf("Failed to set key: %v", err)
	}
	t.Log("SET operation test passed")

	// 测试 GET
	val, err := client.Get(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}
	if val != testValue {
		t.Errorf("Expected value %s, got %s", testValue, val)
	}
	t.Log("GET operation test passed")

	// 测试 EXISTS
	exists, err := client.Exists(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to check existence: %v", err)
	}
	if exists != 1 {
		t.Errorf("Expected key to exist, got %d", exists)
	}
	t.Log("EXISTS operation test passed")

	// 测试 DEL
	deleted, err := client.Del(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}
	if deleted != 1 {
		t.Errorf("Expected 1 deleted key, got %d", deleted)
	}
	t.Log("DEL operation test passed")

	// 验证删除后不存在
	exists, _ = client.Exists(ctx, testKey).Result()
	if exists != 0 {
		t.Errorf("Key should not exist after deletion, got %d", exists)
	}
	t.Log("Redis basic operations test completed successfully")
}

// TestRedisExpiration 测试 Redis 过期时间
func TestRedisExpiration(t *testing.T) {
	client := GetTestRedis()
	if client == nil {
		t.Skip("Redis is not available, skipping Redis tests")
		return
	}

	ctx := context.Background()
	testKey := "go_kit_test:expire"
	testValue := "expire_value"

	// 清理测试数据
	client.Del(ctx, testKey)

	// 设置带过期时间的键（2秒）
	err := client.Set(ctx, testKey, testValue, 2*time.Second).Err()
	if err != nil {
		t.Fatalf("Failed to set key with expiration: %v", err)
	}

	// 立即获取应该存在
	val, err := client.Get(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}
	if val != testValue {
		t.Errorf("Expected value %s, got %s", testValue, val)
	}

	// 检查 TTL
	ttl, err := client.TTL(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to get TTL: %v", err)
	}
	if ttl <= 0 || ttl > 2*time.Second {
		t.Errorf("Expected TTL between 0 and 2s, got %v", ttl)
	}
	t.Logf("TTL test passed: %v", ttl)

	// 等待过期
	time.Sleep(3 * time.Second)

	// 过期后应该不存在
	_, err = client.Get(ctx, testKey).Result()
	if err != redis.Nil {
		t.Errorf("Expected key to be expired, got error: %v", err)
	}
	t.Log("Redis expiration test completed successfully")
}

// TestRedisListOperations 测试 Redis 列表操作
func TestRedisListOperations(t *testing.T) {
	client := GetTestRedis()
	if client == nil {
		t.Skip("Redis is not available, skipping Redis tests")
		return
	}

	ctx := context.Background()
	testKey := "go_kit_test:list"

	// 清理测试数据
	client.Del(ctx, testKey)

	// 测试 LPUSH
	err := client.LPush(ctx, testKey, "value1", "value2", "value3").Err()
	if err != nil {
		t.Fatalf("Failed to LPUSH: %v", err)
	}
	t.Log("LPUSH operation test passed")

	// 测试 LLEN
	length, err := client.LLen(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to get list length: %v", err)
	}
	if length != 3 {
		t.Errorf("Expected list length 3, got %d", length)
	}
	t.Log("LLEN operation test passed")

	// 测试 LRANGE
	values, err := client.LRange(ctx, testKey, 0, -1).Result()
	if err != nil {
		t.Fatalf("Failed to LRANGE: %v", err)
	}
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}
	t.Logf("LRANGE operation test passed: %v", values)

	// 测试 RPOP
	val, err := client.RPop(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to RPOP: %v", err)
	}
	if val != "value1" {
		t.Errorf("Expected 'value1', got %s", val)
	}
	t.Log("RPOP operation test passed")

	// 清理
	client.Del(ctx, testKey)
	t.Log("Redis list operations test completed successfully")
}

// TestRedisHashOperations 测试 Redis 哈希操作
func TestRedisHashOperations(t *testing.T) {
	client := GetTestRedis()
	if client == nil {
		t.Skip("Redis is not available, skipping Redis tests")
		return
	}

	ctx := context.Background()
	testKey := "go_kit_test:hash"

	// 清理测试数据
	client.Del(ctx, testKey)

	// 测试 HSET
	err := client.HSet(ctx, testKey, "field1", "value1", "field2", "value2").Err()
	if err != nil {
		t.Fatalf("Failed to HSET: %v", err)
	}
	t.Log("HSET operation test passed")

	// 测试 HGET
	val, err := client.HGet(ctx, testKey, "field1").Result()
	if err != nil {
		t.Fatalf("Failed to HGET: %v", err)
	}
	if val != "value1" {
		t.Errorf("Expected 'value1', got %s", val)
	}
	t.Log("HGET operation test passed")

	// 测试 HGETALL
	hash, err := client.HGetAll(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("Failed to HGETALL: %v", err)
	}
	if len(hash) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(hash))
	}
	if hash["field1"] != "value1" || hash["field2"] != "value2" {
		t.Errorf("Hash values mismatch: %v", hash)
	}
	t.Log("HGETALL operation test passed")

	// 测试 HDEL
	deleted, err := client.HDel(ctx, testKey, "field1").Result()
	if err != nil {
		t.Fatalf("Failed to HDEL: %v", err)
	}
	if deleted != 1 {
		t.Errorf("Expected 1 deleted field, got %d", deleted)
	}
	t.Log("HDEL operation test passed")

	// 清理
	client.Del(ctx, testKey)
	t.Log("Redis hash operations test completed successfully")
}

// TestEtcdConnection 测试 etcd 连接
func TestEtcdConnection(t *testing.T) {
	cfg := GetTestEtcdCfg()
	client, err := cfg.ConnEtcd()
	if err != nil {
		t.Skipf("etcd is not available, skipping etcd tests: %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 测试连接状态
	status, err := client.Status(ctx, cfg.Endpoints[0])
	if err != nil {
		t.Fatalf("Failed to get etcd status: %v", err)
	}
	if status == nil {
		t.Fatal("Status should not be nil")
	}
	t.Log("etcd connection test passed")
}

// TestEtcdBasicOperations 测试 etcd 基本操作
func TestEtcdBasicOperations(t *testing.T) {
	client := GetTestEtcd()
	if client == nil {
		t.Skip("etcd is not available, skipping etcd tests")
		return
	}
	defer client.Close()

	ctx := context.Background()
	testKey := "go_kit_test/basic"
	testValue := "test_value_123"

	// 清理测试数据
	client.Delete(ctx, testKey)

	// 测试 PUT
	putResp, err := client.Put(ctx, testKey, testValue)
	if err != nil {
		t.Fatalf("Failed to put key: %v", err)
	}
	if putResp == nil {
		t.Fatal("Put response should not be nil")
	}
	t.Log("PUT operation test passed")

	// 测试 GET
	getResp, err := client.Get(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}
	if getResp.Count != 1 {
		t.Errorf("Expected 1 key, got %d", getResp.Count)
	}
	if len(getResp.Kvs) == 0 {
		t.Fatal("Expected at least one key-value pair")
	}
	if string(getResp.Kvs[0].Value) != testValue {
		t.Errorf("Expected value %s, got %s", testValue, string(getResp.Kvs[0].Value))
	}
	t.Log("GET operation test passed")

	// 测试 DELETE
	delResp, err := client.Delete(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}
	if delResp.Deleted != 1 {
		t.Errorf("Expected 1 deleted key, got %d", delResp.Deleted)
	}
	t.Log("DELETE operation test passed")

	// 验证删除后不存在
	getResp, _ = client.Get(ctx, testKey)
	if getResp.Count != 0 {
		t.Errorf("Key should not exist after deletion, got count %d", getResp.Count)
	}
	t.Log("etcd basic operations test completed successfully")
}

// TestEtcdLease 测试 etcd 租约
func TestEtcdLease(t *testing.T) {
	client := GetTestEtcd()
	if client == nil {
		t.Skip("etcd is not available, skipping etcd tests")
		return
	}
	defer client.Close()

	ctx := context.Background()
	testKey := "go_kit_test/lease"
	testValue := "lease_value"

	// 清理测试数据
	client.Delete(ctx, testKey)

	// 创建租约（5秒）
	leaseResp, err := client.Grant(ctx, 5)
	if err != nil {
		t.Fatalf("Failed to grant lease: %v", err)
	}
	leaseID := leaseResp.ID
	t.Logf("Lease granted with ID: %d", leaseID)

	// 使用租约 PUT
	putResp, err := client.Put(ctx, testKey, testValue, clientv3.WithLease(leaseID))
	if err != nil {
		t.Fatalf("Failed to put key with lease: %v", err)
	}
	if putResp == nil {
		t.Fatal("Put response should not be nil")
	}
	t.Log("PUT with lease test passed")

	// 立即获取应该存在
	getResp, err := client.Get(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}
	if getResp.Count != 1 {
		t.Errorf("Expected 1 key, got %d", getResp.Count)
	}

	// 检查租约时间
	ttlResp, err := client.TimeToLive(ctx, leaseID)
	if err != nil {
		t.Fatalf("Failed to get TTL: %v", err)
	}
	if ttlResp.TTL <= 0 {
		t.Errorf("Expected positive TTL, got %d", ttlResp.TTL)
	}
	t.Logf("TTL test passed: %d seconds", ttlResp.TTL)

	// 等待租约过期
	time.Sleep(6 * time.Second)

	// 过期后应该不存在
	getResp, _ = client.Get(ctx, testKey)
	if getResp.Count != 0 {
		t.Errorf("Key should be expired, got count %d", getResp.Count)
	}
	t.Log("etcd lease test completed successfully")
}

// TestEtcdTransaction 测试 etcd 事务
func TestEtcdTransaction(t *testing.T) {
	client := GetTestEtcd()
	if client == nil {
		t.Skip("etcd is not available, skipping etcd tests")
		return
	}
	defer client.Close()

	ctx := context.Background()
	testKey1 := "go_kit_test/txn1"
	testKey2 := "go_kit_test/txn2"
	testValue1 := "value1"
	testValue2 := "value2"

	// 清理测试数据
	client.Delete(ctx, testKey1, clientv3.WithPrefix())
	client.Delete(ctx, testKey2, clientv3.WithPrefix())

	// 先设置一个键
	_, err := client.Put(ctx, testKey1, testValue1)
	if err != nil {
		t.Fatalf("Failed to put initial key: %v", err)
	}

	// 创建事务：如果 testKey1 存在，则设置 testKey2
	getResp, err := client.Get(ctx, testKey1)
	if err != nil {
		t.Fatalf("Failed to get key for transaction: %v", err)
	}

	txn := client.Txn(ctx)
	txn = txn.If(clientv3.Compare(clientv3.Version(testKey1), ">", 0))
	txn = txn.Then(clientv3.OpPut(testKey2, testValue2))
	txn = txn.Else(clientv3.OpGet(testKey1))

	txnResp, err := txn.Commit()
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
	if !txnResp.Succeeded {
		t.Fatal("Transaction should succeed")
	}
	t.Log("Transaction test passed")

	// 验证 testKey2 被设置
	getResp, err = client.Get(ctx, testKey2)
	if err != nil {
		t.Fatalf("Failed to get key2: %v", err)
	}
	if getResp.Count != 1 {
		t.Errorf("Expected 1 key, got %d", getResp.Count)
	}
	if string(getResp.Kvs[0].Value) != testValue2 {
		t.Errorf("Expected value %s, got %s", testValue2, string(getResp.Kvs[0].Value))
	}
	t.Log("etcd transaction test completed successfully")

	// 清理
	client.Delete(ctx, testKey1)
	client.Delete(ctx, testKey2)
}
