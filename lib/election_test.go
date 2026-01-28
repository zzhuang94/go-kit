package lib

import (
	"context"
	"testing"
	"time"

	"github.com/zzhuang94/go-kit/db"
)

func TestNewElectionWithRedis(t *testing.T) {
	rdb := db.GetTestRedis()
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}

	key := "test_election_redis"
	val := "test_value_1"

	// Clean up
	rdb.Del(ctx, key)

	// Test: create election instance
	election := NewElectionWithRedis(rdb, key, val)
	if election == nil {
		t.Fatal("NewElectionWithRedis should return an election instance")
	}
	defer election.Stop()

	// Wait a bit for election to potentially become master
	time.Sleep(100 * time.Millisecond)

	// Test: should be able to check IsMaster
	isMaster := election.IsMaster()
	_ = isMaster // May or may not be master depending on timing

	// Test: should be able to get master value
	masterVal := election.GetMaster()
	_ = masterVal

	// Cleanup
	election.Stop()
	rdb.Del(ctx, key)
}

func TestNewElectionWithRedis_NilClient(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("NewElectionWithRedis should panic when client is nil")
		}
	}()
	NewElectionWithRedis(nil, "test_key", "test_val")
}

func TestNewElectionWithEtcd(t *testing.T) {
	etcdClient := db.GetTestEtcd()
	ctx := context.Background()
	// Test etcd connection by trying to get a key
	_, err := etcdClient.Get(ctx, "/hzz/test_connection_check")
	if err != nil {
		t.Skipf("Etcd is not available, skipping election tests: %v", err)
	}

	key := "/hzz/test_election_etcd"
	val := "test_value_1"

	// Clean up
	etcdClient.Delete(ctx, key)

	// Test: create election instance
	election := NewElectionWithEtcd(etcdClient, key, val)
	if election == nil {
		t.Fatal("NewElectionWithEtcd should return an election instance")
	}
	defer election.Stop()

	// Wait a bit for election to potentially become master
	time.Sleep(100 * time.Millisecond)

	// Test: should be able to check IsMaster
	isMaster := election.IsMaster()
	_ = isMaster // May or may not be master depending on timing

	// Test: should be able to get master value
	masterVal := election.GetMaster()
	_ = masterVal

	// Cleanup
	election.Stop()
	etcdClient.Delete(ctx, key)
}

func TestNewElectionWithEtcd_NilClient(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("NewElectionWithEtcd should panic when client is nil")
		}
	}()
	NewElectionWithEtcd(nil, "/hzz/test_key", "test_val")
}

func TestElection_IsMaster(t *testing.T) {
	rdb := db.GetTestRedis()
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}

	key := "test_election_ismaster"
	val := "test_value_ismaster"

	// Clean up
	rdb.Del(ctx, key)

	election := NewElectionWithRedis(rdb, key, val)
	defer election.Stop()

	// Wait for election to potentially become master
	time.Sleep(200 * time.Millisecond)

	// Should be able to check IsMaster without error
	isMaster := election.IsMaster()
	if isMaster {
		// If it's master, verify it can get its own value
		masterVal := election.GetMaster()
		if masterVal == "" {
			t.Error("GetMaster should return non-empty value when IsMaster is true")
		}
	}

	// Cleanup
	election.Stop()
	rdb.Del(ctx, key)
}

func TestElection_Release(t *testing.T) {
	rdb := db.GetTestRedis()
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}

	key := "test_election_release"
	val := "test_value_release"

	// Clean up
	rdb.Del(ctx, key)

	election := NewElectionWithRedis(rdb, key, val)
	defer election.Stop()

	// Wait for election to potentially become master
	time.Sleep(200 * time.Millisecond)

	// Release should work without error
	election.Release()

	// After release, should not be master
	if election.IsMaster() {
		t.Error("Election should not be master after Release")
	}

	// Multiple releases should be safe
	election.Release()
	election.Release()

	// Cleanup
	election.Stop()
	rdb.Del(ctx, key)
}

func TestElection_GetMaster(t *testing.T) {
	rdb := db.GetTestRedis()
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}

	key := "test_election_getmaster"
	val := "test_value_getmaster"

	// Clean up
	rdb.Del(ctx, key)

	election := NewElectionWithRedis(rdb, key, val)
	defer election.Stop()

	// Wait for election to potentially become master
	time.Sleep(200 * time.Millisecond)

	masterVal := election.GetMaster()
	// Master value may be empty if not master yet, or contain the value if master
	_ = masterVal

	// Cleanup
	election.Stop()
	rdb.Del(ctx, key)
}

func TestElection_Stop(t *testing.T) {
	rdb := db.GetTestRedis()
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}

	key := "test_election_stop"
	val := "test_value_stop"

	// Clean up
	rdb.Del(ctx, key)

	election := NewElectionWithRedis(rdb, key, val)

	// Wait a bit
	time.Sleep(100 * time.Millisecond)

	// Stop should work without error
	election.Stop()

	// After stop, should not be master
	if election.IsMaster() {
		t.Error("Election should not be master after Stop")
	}

	// Multiple stops should be safe
	election.Stop()
	election.Stop()

	// Cleanup
	rdb.Del(ctx, key)
}

func TestElection_ConcurrentAccess_Redis(t *testing.T) {
	rdb := db.GetTestRedis()
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}

	key := "test_election_concurrent_redis"
	valPrefix := "test_value"

	// Clean up
	rdb.Del(ctx, key)

	// Create multiple election instances
	elections := make([]*Election, 5)
	for i := 0; i < 5; i++ {
		elections[i] = NewElectionWithRedis(rdb, key, valPrefix+string(rune(i+'0')))
		defer elections[i].Stop()
	}

	// Wait for elections to compete
	time.Sleep(500 * time.Millisecond)

	// Only one should be master
	masterCount := 0
	for _, e := range elections {
		if e.IsMaster() {
			masterCount++
		}
	}

	if masterCount != 1 {
		t.Errorf("Expected exactly 1 master, got %d", masterCount)
	}

	// Verify GetMaster returns the correct value
	masterVal := elections[0].GetMaster()
	if masterVal == "" {
		t.Error("GetMaster should return non-empty value when there is a master")
	}

	// Cleanup
	for _, e := range elections {
		e.Stop()
	}
	rdb.Del(ctx, key)
}

func TestElection_ConcurrentAccess_Etcd(t *testing.T) {
	etcdClient := db.GetTestEtcd()
	ctx := context.Background()
	// Test etcd connection by trying to get a key
	_, err := etcdClient.Get(ctx, "/hzz/test_connection_check")
	if err != nil {
		t.Skipf("Etcd is not available, skipping election tests: %v", err)
	}

	key := "/hzz/test_election_concurrent_etcd"
	valPrefix := "test_value"

	// Clean up
	etcdClient.Delete(ctx, key)

	// Create multiple election instances
	elections := make([]*Election, 5)
	for i := 0; i < 5; i++ {
		elections[i] = NewElectionWithEtcd(etcdClient, key, valPrefix+string(rune(i+'0')))
		defer elections[i].Stop()
	}

	// Wait for elections to compete
	time.Sleep(500 * time.Millisecond)

	// Only one should be master
	masterCount := 0
	for _, e := range elections {
		if e.IsMaster() {
			masterCount++
		}
	}

	if masterCount != 1 {
		t.Errorf("Expected exactly 1 master, got %d", masterCount)
	}

	// Verify GetMaster returns the correct value
	masterVal := elections[0].GetMaster()
	if masterVal == "" {
		t.Error("GetMaster should return non-empty value when there is a master")
	}

	// Cleanup
	for _, e := range elections {
		e.Stop()
	}
	etcdClient.Delete(ctx, key)
}

func TestElection_MasterFailover_Redis(t *testing.T) {
	rdb := db.GetTestRedis()
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}

	key := "test_election_failover_redis"
	val1 := "test_value_1"
	val2 := "test_value_2"

	// Clean up
	rdb.Del(ctx, key)

	// Create first election
	election1 := NewElectionWithRedis(rdb, key, val1)
	defer election1.Stop()

	// Wait for it to become master
	time.Sleep(300 * time.Millisecond)

	if !election1.IsMaster() {
		t.Fatal("First election should become master")
	}

	// Create second election
	election2 := NewElectionWithRedis(rdb, key, val2)
	defer election2.Stop()

	// Wait a bit
	time.Sleep(100 * time.Millisecond)

	// First should still be master (it's renewing its lease)
	if !election1.IsMaster() {
		t.Error("First election should still be master")
	}
	if election2.IsMaster() {
		t.Error("Second election should not be master while first is active")
	}

	// Stop first election
	election1.Stop()

	// Wait for failover
	time.Sleep(500 * time.Millisecond)

	// Second should become master
	if !election2.IsMaster() {
		t.Error("Second election should become master after first stops")
	}

	// Cleanup
	election2.Stop()
	rdb.Del(ctx, key)
}

func TestElection_MasterFailover_Etcd(t *testing.T) {
	etcdClient := db.GetTestEtcd()
	ctx := context.Background()
	// Test etcd connection by trying to get a key
	_, err := etcdClient.Get(ctx, "/hzz/test_connection_check")
	if err != nil {
		t.Skipf("Etcd is not available, skipping election tests: %v", err)
	}

	key := "/hzz/test_election_failover_etcd"
	val1 := "test_value_1"
	val2 := "test_value_2"

	// Clean up
	etcdClient.Delete(ctx, key)

	// Create first election
	election1 := NewElectionWithEtcd(etcdClient, key, val1)
	defer election1.Stop()

	// Wait for it to become master
	time.Sleep(300 * time.Millisecond)

	if !election1.IsMaster() {
		t.Fatal("First election should become master")
	}

	// Create second election
	election2 := NewElectionWithEtcd(etcdClient, key, val2)
	defer election2.Stop()

	// Wait a bit
	time.Sleep(100 * time.Millisecond)

	// First should still be master (it's renewing its lease)
	if !election1.IsMaster() {
		t.Error("First election should still be master")
	}
	if election2.IsMaster() {
		t.Error("Second election should not be master while first is active")
	}

	// Stop first election
	election1.Stop()

	// Wait for failover
	time.Sleep(500 * time.Millisecond)

	// Second should become master
	if !election2.IsMaster() {
		t.Error("Second election should become master after first stops")
	}

	// Cleanup
	election2.Stop()
	etcdClient.Delete(ctx, key)
}
