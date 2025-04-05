package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	counter map[string]int
	lock    sync.RWMutex
}

func (c *Counter) Increment(str string) {
	c.lock.Lock()
	c.counter[str]++
	c.lock.Unlock()
}

func (c *Counter) Value(str string) int {
	defer c.lock.RUnlock()
	c.lock.RLock()
	return c.counter[str]
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.RLock()
	for k, v := range c.counter {
		fn(k, v)
	}
	c.lock.RUnlock()
}

func NewCounter() *Counter {
	return &Counter{counter: make(map[string]int)}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
