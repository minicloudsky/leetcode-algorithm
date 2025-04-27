package redis_distribute_lock_conflict

import (
	"github.com/go-redis/redis/v8"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 模拟资源竞争的并发场景
func TestLockConflict(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	rdb := redis.NewClient(&redis.Options{Addr: redisAddress})

	var wg sync.WaitGroup
	// 启动 3 个并发客户端竞争锁
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			processWithLock(rdb, clientID)
		}(i)
	}
	wg.Wait()
}
