package redis_distribute_lock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

type RedisLock struct {
	client     *redis.Client
	key        string
	value      string
	expiration time.Duration
	unlockCh   chan struct{}
}

func NewRedisLock(client *redis.Client, key string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		client:     client,
		key:        key,
		value:      generateUniqueValue(),
		expiration: expiration,
		unlockCh:   make(chan struct{}),
	}
}

// Lock 加锁（含自动续期）
func (l *RedisLock) Lock(ctx context.Context) (bool, error) {
	ok, err := l.client.SetNX(ctx, l.key, l.value, l.expiration).Result()
	if err != nil || !ok {
		return false, err
	}

	// 启动看门狗自动续期
	go l.watchDog(ctx)
	return true, nil
}

// 自动续期
func (l *RedisLock) watchDog(ctx context.Context) {
	ticker := time.NewTicker(l.expiration / 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 续期：仅当锁仍属于当前客户端时延长过期时间
			success, err := l.client.Eval(ctx, `
				if redis.call("GET", KEYS[1]) == ARGV[1] then
					return redis.call("PEXPIRE", KEYS[1], ARGV[2])
				else
					return 0
				end
			`, []string{l.key}, l.value, l.expiration.Milliseconds()).Bool()

			if err != nil || !success {
				return
			}
		case <-l.unlockCh:
			return
		}
	}
}

// Unlock 解锁
func (l *RedisLock) Unlock(ctx context.Context) error {
	defer close(l.unlockCh)

	script := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`

	result, err := l.client.Eval(ctx, script, []string{l.key}, l.value).Result()
	if err != nil {
		return err
	}
	if result == 0 {
		return fmt.Errorf("unlock failed: lock not exists or value mismatch")
	}
	return nil
}

func generateUniqueValue() string {
	rand.New(rand.NewSource(100000)).Int()
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int())
}
