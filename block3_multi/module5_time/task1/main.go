package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue struct {
	//myData    []int
	/*maxLength int
	blockChan chan bool*/
	myData chan int
}

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	if block {
		val := <-q.myData
		return val, nil
	}
	select {
	case val := <-q.myData:
		return val, nil
	default:
		return 0, ErrEmpty
	}

}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	if block {
		q.myData <- val
		return nil
	}
	select {
	case q.myData <- val:
		return nil
	default:
		return ErrFull
	}

}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	return Queue{myData: make(chan int, n)}
}

// конец решения

func main() {
	q := MakeQueue(2)

	go func() {
		err := q.Put(11, true)
		fmt.Println("put 11:", err)
		// put 11: <nil>

		err = q.Put(12, true)
		fmt.Println("put 12:", err)
		// put 12: <nil>
	}()

	res, err := q.Get(true)
	fmt.Println("get:", res, err)
	// get: 11 <nil>

	res, err = q.Get(true)
	fmt.Println("get:", res, err)
}
