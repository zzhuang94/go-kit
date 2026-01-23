package lib

import (
	"context"
	"strings"
	"testing"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func TestTryCheckIn_Local(t *testing.T) {
	// Test local limiter (nil redis client)
	limiter, err := TryCheckIn(nil, "test_limiter_local", 5, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed: %v", err)
	}
	if limiter == nil {
		t.Fatal("TryCheckIn should return a limiter")
	}
	
	// Test release
	limiter.Release()
}

func TestTryCheckIn_LocalLimit(t *testing.T) {
	// Test that limit is enforced
	limit := 3
	
	// Acquire limit number of limiters
	limiters := make([]*limiter, limit)
	for i := 0; i < limit; i++ {
		l, err := TryCheckIn(nil, "test_limiter_limit", limit, 1*time.Second)
		if err != nil {
			t.Fatalf("TryCheckIn failed at %d: %v", i, err)
		}
		if l == nil {
			t.Fatalf("TryCheckIn returned nil at %d", i)
		}
		limiters[i] = l
	}
	
	// Try to exceed limit (should timeout)
	start := time.Now()
	l, err := TryCheckIn(nil, "test_limiter_limit", limit, 100*time.Millisecond)
	duration := time.Since(start)
	
	if err == nil {
		t.Error("TryCheckIn should timeout when limit is reached")
		if l != nil {
			l.Release()
		}
	}
	if duration < 50*time.Millisecond || duration > 200*time.Millisecond {
		t.Logf("Timeout duration: %v (expected around 100ms)", duration)
	}
	
	// Release all limiters
	for _, l := range limiters {
		l.Release()
	}
}

func TestTryCheckIn_LocalRelease(t *testing.T) {
	// Test that release allows new check-ins
	limit := 2
	
	// Acquire limit
	limiter1, err := TryCheckIn(nil, "test_limiter_release", limit, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed: %v", err)
	}
	
	limiter2, err := TryCheckIn(nil, "test_limiter_release", limit, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed: %v", err)
	}
	
	// Release one
	limiter1.Release()
	
	// Should be able to acquire again
	limiter3, err := TryCheckIn(nil, "test_limiter_release", limit, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed after release: %v", err)
	}
	if limiter3 == nil {
		t.Fatal("TryCheckIn should return a limiter after release")
	}
	
	limiter2.Release()
	limiter3.Release()
}

func TestTryCheckIn_LocalMultipleRelease(t *testing.T) {
	limiter, err := TryCheckIn(nil, "test_limiter_multirelease", 5, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed: %v", err)
	}
	
	// Multiple releases should be safe
	limiter.Release()
	limiter.Release()
	limiter.Release()
}

func TestTryCheckIn_LocalDifferentKeys(t *testing.T) {
	// Test that different keys don't interfere
	limiter1, err := TryCheckIn(nil, "test_limiter_key1", 2, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed for key1: %v", err)
	}
	
	limiter2, err := TryCheckIn(nil, "test_limiter_key2", 2, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed for key2: %v", err)
	}
	
	limiter1.Release()
	limiter2.Release()
}

func TestTryCheckIn_MaxLimit(t *testing.T) {
	// Test with maximum limit
	maxLimit := int(^uint(0) >> 1)
	
	limiter, err := TryCheckIn(nil, "test_limiter_max", maxLimit, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed with max limit: %v", err)
	}
	if limiter == nil {
		t.Fatal("TryCheckIn should return a limiter")
	}
	limiter.Release()
}

func TestTryCheckIn_InvalidLimit(t *testing.T) {
	// Test with limit exceeding max
	maxLimit := int(^uint(0) >> 1)
	invalidLimit := maxLimit + 1
	
	limiter, err := TryCheckIn(nil, "test_limiter_invalid", invalidLimit, 1*time.Second)
	if err == nil {
		t.Error("TryCheckIn should fail with invalid limit")
		if limiter != nil {
			limiter.Release()
		}
	}
	if err != nil && !strings.Contains(err.Error(), "限制数不得大于") {
		t.Logf("Expected error about limit, got: %v", err)
	}
}

func TestTryCheckIn_Redis(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping Redis tests")
	}
	
	// Clean up any existing limiter key
	rdb.Del(ctx, "test_limiter_redis")
	
	// Test Redis limiter
	limiter, err := TryCheckIn(rdb, "test_limiter_redis", 5, 1*time.Second)
	if err != nil {
		t.Fatalf("TryCheckIn failed: %v", err)
	}
	if limiter == nil {
		t.Fatal("TryCheckIn should return a limiter")
	}
	
	// Test release
	limiter.Release()
	
	// Cleanup
	rdb.Del(ctx, "test_limiter_redis")
}

func TestTryCheckIn_RedisLimit(t *testing.T) {
	// Skip if Redis is not available
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		t.Skip("Redis is not available, skipping Redis tests")
	}
	
	// Clean up
	rdb.Del(ctx, "test_limiter_redis_limit")
	
	limit := 3
	
	// Acquire limit number of limiters
	limiters := make([]*limiter, limit)
	for i := 0; i < limit; i++ {
		l, err := TryCheckIn(rdb, "test_limiter_redis_limit", limit, 1*time.Second)
		if err != nil {
			t.Fatalf("TryCheckIn failed at %d: %v", i, err)
		}
		if l == nil {
			t.Fatalf("TryCheckIn returned nil at %d", i)
		}
		limiters[i] = l
	}
	
	// Try to exceed limit (should timeout)
	start := time.Now()
	l, err := TryCheckIn(rdb, "test_limiter_redis_limit", limit, 100*time.Millisecond)
	duration := time.Since(start)
	
	if err == nil {
		t.Error("TryCheckIn should timeout when limit is reached")
		if l != nil {
			l.Release()
		}
	}
	if duration < 50*time.Millisecond || duration > 200*time.Millisecond {
		t.Logf("Timeout duration: %v (expected around 100ms)", duration)
	}
	
	// Release all limiters
	for _, l := range limiters {
		l.Release()
	}
	
	// Cleanup
	rdb.Del(ctx, "test_limiter_redis_limit")
}
