package redis_distribute_lock_conflict

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

const (
	lockKey      = "concurrent_resource_lock"
	redisAddress = "localhost:6379"
)

// RedisLock 改进版分布式锁（含冲突处理）
type RedisLock struct {
	client        *redis.Client
	lockKey       string
	uniqueValue   string
	expiration    time.Duration
	retryInterval time.Duration
	maxRetries    int
}

func NewRedisLock(client *redis.Client, lockKey string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		client:        client,
		lockKey:       lockKey,
		uniqueValue:   generateUniqueValue(),
		expiration:    expiration,
		retryInterval: 100 * time.Millisecond,
		maxRetries:    5,
	}
}

func processWithLock(rdb *redis.Client, clientID int) {
	lock := NewRedisLock(rdb, lockKey, 5*time.Second)
	ctx := context.Background()

	// 带重试机制的加锁
	acquired := false
	for attempt := 1; attempt <= lock.maxRetries; attempt++ {
		ok, err := lock.TryLock(ctx)
		if err != nil {
			fmt.Printf("Client %d attempt %d: Lock error - %v\n", clientID, attempt, err)
			continue
		}
		if ok {
			acquired = true
			break
		}
		fmt.Printf("Client %d attempt %d: Lock occupied, retrying...\n", clientID, attempt)
		time.Sleep(lock.retryInterval)
	}

	if !acquired {
		fmt.Printf("Client %d: Failed to acquire lock after %d attempts\n", clientID, lock.maxRetries)
		return
	}

	defer func() {
		if err := lock.Unlock(ctx); err != nil {
			fmt.Printf("Client %d unlock error: %v\n", clientID, err)
		}
	}()

	// 模拟业务处理
	fmt.Printf("Client %d acquired lock, processing...\n", clientID)
	processTime := time.Duration(2+rand.Intn(3)) * time.Second
	time.Sleep(processTime)
	fmt.Printf("Client %d completed processing\n", clientID)
}

// TryLock 带冲突检测的加锁方法
func (l *RedisLock) TryLock(ctx context.Context) (bool, error) {
	result, err := l.client.SetNX(ctx, l.lockKey, l.uniqueValue, l.expiration).Result()
	if err != nil {
		return false, fmt.Errorf("redis setnx error: %w", err)
	}

	// 加锁成功时启动自动续期
	if result {
		go l.autoRenew(ctx)
	}
	return result, nil
}

// 自动续期机制
func (l *RedisLock) autoRenew(ctx context.Context) {
	ticker := time.NewTicker(l.expiration / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 使用 Lua 脚本保证原子性续期
			success, err := l.client.Eval(ctx, `
				if redis.call("GET", KEYS[1]) == ARGV[1] then
					return redis.call("PEXPIRE", KEYS[1], ARGV[2])
				else
					return 0
				end
			`, []string{l.lockKey}, l.uniqueValue, l.expiration.Milliseconds()).Bool()

			if err != nil || !success {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

// Unlock 安全解锁
func (l *RedisLock) Unlock(ctx context.Context) error {
	script := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`
	result, err := l.client.Eval(ctx, script, []string{l.lockKey}, l.uniqueValue).Int64()
	if err != nil {
		return fmt.Errorf("unlock script failed: %w", err)
	}
	if result == 0 {
		return fmt.Errorf("unlock failed: lock not exists or value mismatch")
	}
	return nil
}

func generateUniqueValue() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
}
