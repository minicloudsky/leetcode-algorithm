package main

import (
	"fmt"
	"time"
)

// 使用 select 语句来轮询 a、b 和 c。
// 在 select 中，如果 a 有数据，就优先选择 a，否则检查 b 和 c。
// 使用 time.After 来避免阻塞，避免死锁的情况

func priorityChannel() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		a <- 1 // 模拟a有数据，延迟1秒
	}()

	go func() {
		time.Sleep(2 * time.Second)
		b <- 2 // 模拟b有数据，延迟2秒
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c <- 3 // 模拟c有数据，延迟3秒
	}()

	for {
		select {
		case msgA := <-a:
			fmt.Println("Received from a:", msgA)
		case msgB := <-b:
			fmt.Println("Received from b:", msgB)
		case msgC := <-c:
			fmt.Println("Received from c:", msgC)
		}
	}
}
func priorityChannel2() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		a <- 1 // 模拟a有数据，延迟1秒
	}()

	go func() {
		time.Sleep(2 * time.Second)
		b <- 2 // 模拟b有数据，延迟2秒
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c <- 3 // 模拟c有数据，延迟3秒
	}()

	for {
		select {
		case msgA := <-a:
			fmt.Println("Received from a:", msgA)
		case msgB := <-b:
			fmt.Println("Received from b:", msgB)
		case msgC := <-c:
			fmt.Println("Received from c:", msgC)
		}
	}
}
func main() {
	priorityChannel()
	priorityChannel2()
}
