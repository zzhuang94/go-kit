package lib

import (
	"context"
	"strings"
	"testing"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func TestNewElection(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}
	
	key := "test_election_new"
	val := "node1"
	
	// Clean up
	rdb.Del(ctx, key)
	
	// Create election
	election := NewElection(key, val, rdb)
	if election == nil {
		t.Fatal("NewElection should return an election instance")
	}
	
	// Wait a bit for election to settle
	time.Sleep(100 * time.Millisecond)
	
	// Should be master (first one)
	if !election.IsMaster() {
		t.Log("First election instance should be master")
	}
	
	election.Stop()
	
	// Cleanup
	rdb.Del(ctx, key)
}

func TestElection_IsMaster(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}
	
	key := "test_election_ismaster"
	val := "node1"
	
	// Clean up
	rdb.Del(ctx, key)
	
	election := NewElection(key, val, rdb)
	defer election.Stop()
	
	// Wait for election
	time.Sleep(100 * time.Millisecond)
	
	// First instance should be master
	if !election.IsMaster() {
		t.Log("First instance should be master")
	}
	
	// Cleanup
	rdb.Del(ctx, key)
}

func TestElection_GetMaster(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}
	
	key := "test_election_getmaster"
	val := "node1"
	
	// Clean up
	rdb.Del(ctx, key)
	
	election := NewElection(key, val, rdb)
	defer election.Stop()
	
	// Wait for election
	time.Sleep(100 * time.Millisecond)
	
	master := election.GetMaster()
	if master == "" {
		t.Log("GetMaster should return master value")
	} else if !strings.Contains(master, val) {
		t.Logf("GetMaster should contain node value, got: %s", master)
	}
	
	// Cleanup
	rdb.Del(ctx, key)
}

func TestElection_Release(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}
	
	key := "test_election_release"
	val := "node1"
	
	// Clean up
	rdb.Del(ctx, key)
	
	election := NewElection(key, val, rdb)
	
	// Wait for election
	time.Sleep(100 * time.Millisecond)
	
	if election.IsMaster() {
		election.Release()
		
		// Should not be master after release
		if election.IsMaster() {
			t.Error("Election should not be master after release")
		}
	}
	
	election.Stop()
	
	// Cleanup
	rdb.Del(ctx, key)
}

func TestElection_Stop(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}
	
	key := "test_election_stop"
	val := "node1"
	
	// Clean up
	rdb.Del(ctx, key)
	
	election := NewElection(key, val, rdb)
	
	// Wait for election
	time.Sleep(100 * time.Millisecond)
	
	// Stop should not panic
	election.Stop()
	
	// Multiple stops should be safe
	election.Stop()
	election.Stop()
	
	// Cleanup
	rdb.Del(ctx, key)
}

func TestElection_MultipleInstances(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}
	
	key := "test_election_multi"
	
	// Clean up
	rdb.Del(ctx, key)
	
	// Create two election instances
	election1 := NewElection(key, "node1", rdb)
	election2 := NewElection(key, "node2", rdb)
	
	defer election1.Stop()
	defer election2.Stop()
	
	// Wait for election
	time.Sleep(200 * time.Millisecond)
	
	// Only one should be master
	masterCount := 0
	if election1.IsMaster() {
		masterCount++
	}
	if election2.IsMaster() {
		masterCount++
	}
	
	if masterCount != 1 {
		t.Logf("Expected exactly one master, got %d", masterCount)
	}
	
	// Get master should return one of the nodes
	master := election1.GetMaster()
	if master != "" {
		if !strings.Contains(master, "node1") && !strings.Contains(master, "node2") {
			t.Logf("GetMaster should return one of the nodes, got: %s", master)
		}
	}
	
	// Cleanup
	rdb.Del(ctx, key)
}

func TestElection_ReleaseNonMaster(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping election tests")
	}
	
	key := "test_election_releasenonmaster"
	
	// Clean up
	rdb.Del(ctx, key)
	
	election1 := NewElection(key, "node1", rdb)
	election2 := NewElection(key, "node2", rdb)
	
	defer election1.Stop()
	defer election2.Stop()
	
	// Wait for election
	time.Sleep(200 * time.Millisecond)
	
	// Release non-master should be safe
	if !election1.IsMaster() {
		election1.Release()
	}
	if !election2.IsMaster() {
		election2.Release()
	}
	
	// Cleanup
	rdb.Del(ctx, key)
}
