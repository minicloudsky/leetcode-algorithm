package main

import (
	"fmt"
	"runtime"
	"sync"
)

func producer(c chan int, val int, errChan chan error) {
	c <- val
}

func consumer(c chan int, errChan chan error, wg *sync.WaitGroup) {
	select {
	case val, ok := <-c:
		if ok {
			fmt.Println(val)
		}
	case err := <-errChan:
		fmt.Printf("err: %v received, return", err)
		return
	}
}

func main() {
	c := make(chan int, 100)
	errChan := make(chan error, 1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go producer(c, i, errChan)
		fmt.Println("channel size: ", len(c))
		wg.Add(1)
		go consumer(c, errChan, wg)
	}
	errChan <- fmt.Errorf("send finished")
	wg.Wait()
	runtime.NumGoroutine()
}
