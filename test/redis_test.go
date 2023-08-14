package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

func TestSetAndGet(t *testing.T) {
	key := "test_key"
	value := "test_value"

	err := rdb.Set(ctx, key, value, 10*time.Second).Err()
	if err != nil {
		t.Errorf("Failed to set value: %v", err)
		return
	}

	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		t.Errorf("Failed to get value: %v", err)
		return
	}

	if result != value {
		t.Errorf("Expected %s, got %s", value, result)
	}
}

func TestSetWithExpiration(t *testing.T) {
	key := "expiring_key"
	value := "expiring_value"
	expiration := 5 * time.Second

	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		t.Errorf("Failed to set value with expiration: %v", err)
		return
	}

	time.Sleep(expiration + time.Second)
	_, err = rdb.Get(ctx, key).Result()
	if err == nil || err != redis.Nil {
		t.Errorf("Expected value to be expired, but it was retrieved")
	}
}

func TestNonExistentKey(t *testing.T) {
	nonExistentKey := "non_existent_key"

	_, err := rdb.Get(ctx, nonExistentKey).Result()
	if err == nil || err != redis.Nil {
		t.Errorf("Expected redis.Nil error for non-existent key, got: %v", err)
	}
}

func TestMain(m *testing.M) {
	code := m.Run()

	if err := rdb.Close(); err != nil {
		fmt.Println("Failed to close Redis connection:", err)
	}

	os.Exit(code)
}
