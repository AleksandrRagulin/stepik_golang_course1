package main

import "fmt"

// IntSet реализует множество целых чисел
// (элементы множества уникальны).
/*type IntSet struct {
}*/
type IntSet map[int]bool

// MakeIntSet создает пустое множество.
func MakeIntSet() IntSet {
	return make(IntSet)
}

// Contains проверяет, содержится ли элемент в множестве.
func (s IntSet) Contains(elem int) bool {
	_, ok := s[elem]
	return ok
}

// Add добавляет элемент в множество.
// Возвращает true, если элемент добавлен,
// иначе false (если элемент уже содержится в множестве).
func (s IntSet) Add(elem int) bool {
	_, ok := s[elem]
	if ok {
		return false
	}
	s[elem] = true
	return true
}

func main() {

	fmt.Println("go is awesome")
	// true
}
