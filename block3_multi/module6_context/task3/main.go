package main

import (
	"errors"
	"fmt"
	"time"
)

// начало решения

// ErrFailed и ErrManual - причины остановки цикла.
var ErrFailed = errors.New("failed")
var ErrManual = errors.New("manual")

// Worker выполняет заданную функцию в цикле, пока не будет остановлен.
// Гарантируется, что Worker используется только в одной горутине.
type Worker struct {
	fn        func() error
	myErr     error
	start     bool
	stop      bool
	afterFunc []func()
	// TODO: добавить поля
}

// NewWorker создает новый экземпляр Worker с заданной функцией.
// Но пока не запускает цикл с функцией.
func NewWorker(fn func() error) *Worker {
	return &Worker{fn: fn, myErr: nil, start: false, stop: false}
}

// Start запускает отдельную горутину, в которой циклически
// выполняет заданную функцию, пока не будет вызван метод Stop,
// либо пока функция не вернет ошибку.
// Повторные вызовы Start игнорируются.
func (w *Worker) Start() {
	// TODO: реализовать требования
	if w.stop || w.start {
		return
	}

	w.start = true

	go func() {
		for {
			if w.stop {
				for _, fn := range w.afterFunc {
					fn()
				}
				return
			}

			err := w.fn()
			if err != nil {
				w.myErr = ErrFailed
				w.stop = true
				w.start = false

				for _, fn := range w.afterFunc {
					fn()
				}
				return
			}
		}
	}()
}

// Stop останавливает выполнение цикла.
// Вызов Stop до Start игнорируется.
// Повторные вызовы Stop игнорируются.
func (w *Worker) Stop() {
	// TODO: реализовать требования
	if w.stop {
		return
	}

	w.myErr = ErrManual
	w.stop = true
	w.start = false

}

// AfterStop регистрирует функцию, которая
// будет вызвана после остановки цикла.
// Можно зарегистрировать несколько функций.
// Вызовы AfterStop после Start игнорируются.
func (w *Worker) AfterStop(fn func()) {
	// TODO: реализовать требования
	if w.start {
		return
	}

	w.afterFunc = append(w.afterFunc, fn)

}

// Err возвращает причину остановки цикла:
// - ErrManual - вручную через метод Stop;
// - ErrFailed - из-за ошибки, которую вернула функция.
func (w *Worker) Err() error {
	// TODO: реализовать требования
	return w.myErr
}

// конец решения

func main() {
	{
		// Start-Stop
		count := 9
		fn := func() error {
			fmt.Print(count, " ")
			count--
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := NewWorker(fn)
		worker.Start()
		time.Sleep(105 * time.Millisecond)
		worker.Stop()

		fmt.Println()
		// 9 8 7 6 5 4 3 2 1 0
	}
	{
		// ErrFailed
		count := 3
		fn := func() error {
			fmt.Print(count, " ")
			count--
			if count == 0 {
				return errors.New("count is zero")
			}
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := NewWorker(fn)
		worker.Start()
		time.Sleep(35 * time.Millisecond)
		worker.Stop()

		fmt.Println(worker.Err())
		// 3 2 1 failed
	}
	{
		// AfterStop
		fn := func() error { return nil }

		worker := NewWorker(fn)
		worker.AfterStop(func() {
			fmt.Println("called after stop")
		})

		worker.Start()
		worker.Stop()

		time.Sleep(10 * time.Millisecond)
		// called after stop
	}
}
