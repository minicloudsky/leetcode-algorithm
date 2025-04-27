package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

// 使用示例
func TestRedLock(t *testing.T) {
	// 初始化 5 个 Redis 节点（生产环境应为独立实例）
	clients := []*redis.Client{
		redis.NewClient(&redis.Options{Addr: "node1:6379"}),
		redis.NewClient(&redis.Options{Addr: "node2:6379"}),
		redis.NewClient(&redis.Options{Addr: "node3:6379"}),
		redis.NewClient(&redis.Options{Addr: "node4:6379"}),
		redis.NewClient(&redis.Options{Addr: "node5:6379"}),
	}

	redLock := NewRedLock(clients, "critical_section", 10*time.Second)
	ctx := context.Background()

	// 尝试获取锁
	locked, err := redLock.Lock(ctx)
	if err != nil {
		panic(err)
	}
	if !locked {
		fmt.Println("Failed to acquire lock")
		return
	}
	defer func(redLock *RedLock, ctx context.Context) {
		err := redLock.Unlock(ctx)
		if err != nil {
			return
		}
	}(redLock, ctx)

	// 执行业务逻辑
	fmt.Println("Lock acquired, processing...")
	time.Sleep(8 * time.Second)
}
