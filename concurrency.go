package main

import (
	"fmt"
	"sync"
)

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibo(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

type Counter struct {
	mu sync.Mutex
	v  map[string]int
}

// using mutexes
func (c *Counter) inc(key string) {
	c.mu.Lock()

	c.v[key]++

	c.mu.Unlock()
}

func (c *Counter) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func con() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// receiving from a channel is a blocking call
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
