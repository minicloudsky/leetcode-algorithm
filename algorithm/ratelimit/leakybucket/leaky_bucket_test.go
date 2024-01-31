package leakybucket

import (
	"fmt"
	"sync"
	"testing"
)

func TestLeakyBucket(t *testing.T) {
	l := LeakyBucket{lock: sync.Mutex{}}
	l.Set(100, 100)
	for i := 0; i < 1000; i++ {
		if l.Allow() {
			fmt.Println("正常请求")
		} else {
			fmt.Println("limited!")
		}
	}
}
