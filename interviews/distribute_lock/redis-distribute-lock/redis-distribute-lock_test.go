package redis_distribute_lock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"testing"
	"time"
)

func TestRedisDistributeLock(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	resourceKey := fmt.Sprintf("createOrder_%v", rand.Int())

	lock := NewRedisLock(rdb, resourceKey, 30*time.Second)
	ctx := context.Background()

	// 尝试加锁
	acquired, err := lock.Lock(ctx)
	if err != nil {
		panic(err)
	}
	if !acquired {
		fmt.Println("Failed to acquire lock")
		return
	}
	defer func(lock *RedisLock, ctx context.Context) {
		err := lock.Unlock(ctx)
		if err != nil {
			fmt.Printf("unlock err! %v", err)
		}
		fmt.Println("unlock success...")
	}(lock, ctx)

	// 执行业务逻辑
	fmt.Println("Lock acquired, processing...")
	time.Sleep(5 * time.Second)
}

func TestGenerateUniqueValue(t *testing.T) {
	uuid := generateUniqueValue()
	fmt.Println(uuid)
}
