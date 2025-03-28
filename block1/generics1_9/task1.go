package main

import "fmt"

func Produce[T any](val T, n int) []T {
	vals := make([]T, n)
	for i := range n {
		vals[i] = val
	}
	return vals
}

func main() {
	// так работает
	intSlice := Produce(5, 3)
	fmt.Println(intSlice)
	// [5 5 5]

	// а так уже нет
	strSlice := Produce("o", 5)
	fmt.Println(strSlice)
	// [o o o o o]
}
