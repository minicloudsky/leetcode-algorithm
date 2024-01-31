package counter

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	lr := LimitRate{
		begin: time.Now(),
		cycle: 1,
		lock:  sync.Mutex{},
	}
	lr.Set(100, 1*time.Second)
	// 请求过多，触发限流
	for i := 0; i < 1000000; i++ {
		if !lr.Allow() {
			fmt.Println("rate limited,exit")
			break
		}
	}
	// 可以正常请求
	for i := 0; i < 100; i++ {
		if !lr.Allow() {
			fmt.Println("rate limited,exit")
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}
