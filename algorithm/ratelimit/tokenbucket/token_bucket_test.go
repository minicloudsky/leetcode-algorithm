package tokenbucket

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	tb := TokenBucket{lock: sync.Mutex{}}
	// 先积累1s的令牌
	tb.Init(100, 1000)
	time.Sleep(1 * time.Second)
	for i := 0; i < 10000000; i++ {
		if i < 1000 {
			time.Sleep(10 * time.Millisecond)
		}
		if tb.Allow() {
			fmt.Println("正常处理请求")
		} else {
			fmt.Println("限流中...")
			break
		}
	}
}
