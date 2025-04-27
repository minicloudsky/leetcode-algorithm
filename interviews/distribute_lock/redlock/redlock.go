package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"sync"
	"time"
)

// RedLock 配置
type RedLock struct {
	clients    []*redis.Client // 多个独立 Redis 节点
	quorum     int             // 多数节点数（N/2 +1）
	lockKey    string
	uniqueVal  string
	expiration time.Duration
	drift      time.Duration // 时钟漂移补偿
}

func NewRedLock(clients []*redis.Client, lockKey string, expiration time.Duration) *RedLock {
	return &RedLock{
		clients:    clients,
		quorum:     len(clients)/2 + 1,
		lockKey:    lockKey,
		uniqueVal:  generateUniqueValue(),
		expiration: expiration,
		drift:      2 * time.Second, // 典型时钟漂移补偿
	}
}

// 尝试获取锁（核心逻辑）
func (rl *RedLock) Lock(ctx context.Context) (bool, error) {
	startTime := time.Now()

	var wg sync.WaitGroup
	successCh := make(chan bool, len(rl.clients))
	errCh := make(chan error, len(rl.clients))

	// 并发向所有节点请求锁
	for _, client := range rl.clients {
		wg.Add(1)
		go func(c *redis.Client) {
			defer wg.Done()
			ok, err := rl.tryLockNode(ctx, c)
			if err != nil {
				errCh <- err
				return
			}
			successCh <- ok
		}(client)
	}

	// 收集结果
	wg.Wait()
	close(successCh)
	close(errCh)

	// 统计成功数
	successCount := 0
	for ok := range successCh {
		if ok {
			successCount++
		}
	}

	// 计算总耗时
	totalTime := time.Since(startTime)
	validityTime := rl.expiration - totalTime - rl.drift

	// 判断是否满足多数原则且未超时
	if successCount >= rl.quorum && validityTime > 0 {
		go rl.startWatchDog(ctx) // 启动自动续期
		return true, nil
	}

	// 加锁失败时释放已获得的锁
	go rl.rollbackUnlock(ctx)
	return false, nil
}

// 单节点加锁
func (rl *RedLock) tryLockNode(ctx context.Context, client *redis.Client) (bool, error) {
	result, err := client.SetNX(ctx, rl.lockKey, rl.uniqueVal, rl.expiration).Result()
	if err == redis.Nil {
		return false, nil // 锁已被占用
	}
	return result, err
}

// 自动续期（看门狗）
func (rl *RedLock) startWatchDog(ctx context.Context) {
	ticker := time.NewTicker(rl.expiration / 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 向所有节点续期
			successCount := 0
			for _, client := range rl.clients {
				ok, _ := rl.renewLock(ctx, client)
				if ok {
					successCount++
				}
			}
			if successCount < rl.quorum {
				return // 不再满足多数原则，停止续期
			}
		case <-ctx.Done():
			return
		}
	}
}

// 单节点续期
func (rl *RedLock) renewLock(ctx context.Context, client *redis.Client) (bool, error) {
	script := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("PEXPIRE", KEYS[1], ARGV[2])
		else
			return 0
		end
	`
	return client.Eval(ctx, script, []string{rl.lockKey}, rl.uniqueVal, rl.expiration.Milliseconds()).Bool()
}

// 解锁所有节点
func (rl *RedLock) Unlock(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(rl.clients))

	for _, client := range rl.clients {
		wg.Add(1)
		go func(c *redis.Client) {
			defer wg.Done()
			if err := rl.unlockNode(ctx, c); err != nil {
				errCh <- err
			}
		}(client)
	}

	wg.Wait()
	close(errCh)

	// 返回首个错误（实际生产需更细致处理）
	if len(errCh) > 0 {
		return <-errCh
	}
	return nil
}

// 单节点解锁
func (rl *RedLock) unlockNode(ctx context.Context, client *redis.Client) error {
	script := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`
	_, err := client.Eval(ctx, script, []string{rl.lockKey}, rl.uniqueVal).Result()
	return err
}

// 加锁失败时回滚已获得的锁
func (rl *RedLock) rollbackUnlock(ctx context.Context) {
	for _, client := range rl.clients {
		rl.unlockNode(ctx, client)
	}
}

func generateUniqueValue() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
}
