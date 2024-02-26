package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func Producer(c chan int, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 0
	for {
		c <- counter
		fmt.Printf("send %d\n", counter)
		time.Sleep(1 * time.Second)
		if counter == 10 {
			errChan <- errors.New("send failed, stop")
			close(c)
			close(errChan)
			return
		}
		counter++
	}
}

func Consumer(c chan int, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case value, ok := <-c:
			if ok {
				fmt.Printf("received %d\n", value)
			}
		case err, ok := <-errChan:
			if ok {
				fmt.Printf("err %v\n", err)
				return
			}
		}
	}
}

func Scheduler() {
	c := make(chan int, 1)
	errChan := make(chan error, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go Producer(c, errChan, &wg)
	go Consumer(c, errChan, &wg)
	wg.Wait()
}

func main() {
	Scheduler()
}
