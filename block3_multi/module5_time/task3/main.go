package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	cancel := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)

		select {
		case <-time.After(dur):
			select {
			case <-cancel:
				return
			default:
				fn()
			}
		case <-cancel:
			return
		}
	}()

	return func() {
		select {
		case cancel <- struct{}{}:
			<-done
		case <-done:
		}
	}
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
