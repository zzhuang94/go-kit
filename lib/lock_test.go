package lib

import (
	"context"
	"testing"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func TestTryLock_Local(t *testing.T) {
	// Test local lock (nil redis client)
	lock, err := TryLock(nil, "test_lock_local", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed: %v", err)
	}
	if lock == nil {
		t.Fatal("TryLock should return a lock")
	}
	
	// Test release
	lock.Release()
	
	// Test that we can acquire lock again after release
	lock2, err := TryLock(nil, "test_lock_local", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed on second attempt: %v", err)
	}
	if lock2 == nil {
		t.Fatal("TryLock should return a lock on second attempt")
	}
	lock2.Release()
}

func TestTryLock_LocalTimeout(t *testing.T) {
	// Test timeout when lock is held
	lock1, err := TryLock(nil, "test_lock_timeout", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed: %v", err)
	}
	
	// Try to acquire same lock (should timeout)
	start := time.Now()
	lock2, err := TryLock(nil, "test_lock_timeout", 100*time.Millisecond)
	duration := time.Since(start)
	
	if err == nil {
		t.Error("TryLock should timeout when lock is held")
		lock2.Release()
	}
	if duration < 50*time.Millisecond || duration > 200*time.Millisecond {
		t.Logf("Timeout duration: %v (expected around 100ms)", duration)
	}
	
	lock1.Release()
}

func TestTryLock_Release(t *testing.T) {
	lock, err := TryLock(nil, "test_lock_release", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed: %v", err)
	}
	
	// Release should not panic
	lock.Release()
	
	// Multiple releases should be safe
	lock.Release()
	lock.Release()
}

func TestTryLock_DifferentKeys(t *testing.T) {
	// Test that different keys don't interfere
	lock1, err := TryLock(nil, "test_lock_key1", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed for key1: %v", err)
	}
	
	lock2, err := TryLock(nil, "test_lock_key2", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed for key2: %v", err)
	}
	
	lock1.Release()
	lock2.Release()
}

func TestTryLock_Redis(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping Redis tests")
	}
	
	// Clean up any existing lock
	rdb.Del(ctx, "test_lock_redis")
	
	// Test Redis lock
	lock, err := TryLock(rdb, "test_lock_redis", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed: %v", err)
	}
	if lock == nil {
		t.Fatal("TryLock should return a lock")
	}
	
	// Test release
	lock.Release()
	
	// Verify lock is released
	val, _ := rdb.Get(ctx, "test_lock_redis").Result()
	if val != "" {
		t.Logf("Lock key still exists after release: %s", val)
	}
	
	// Cleanup
	rdb.Del(ctx, "test_lock_redis")
}

func TestTryLock_RedisTimeout(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping Redis tests")
	}
	
	// Clean up any existing lock
	rdb.Del(ctx, "test_lock_redis_timeout")
	
	// Acquire lock
	lock1, err := TryLock(rdb, "test_lock_redis_timeout", 1*time.Second)
	if err != nil {
		t.Fatalf("TryLock failed: %v", err)
	}
	
	// Try to acquire same lock (should timeout)
	start := time.Now()
	lock2, err := TryLock(rdb, "test_lock_redis_timeout", 100*time.Millisecond)
	duration := time.Since(start)
	
	if err == nil {
		t.Error("TryLock should timeout when lock is held")
		if lock2 != nil {
			lock2.Release()
		}
	}
	if duration < 50*time.Millisecond || duration > 200*time.Millisecond {
		t.Logf("Timeout duration: %v (expected around 100ms)", duration)
	}
	
	lock1.Release()
	
	// Cleanup
	rdb.Del(ctx, "test_lock_redis_timeout")
}
