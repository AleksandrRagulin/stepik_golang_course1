package main

import "fmt"

//https://stepik.org/lesson/1352836/step/11?unit=1368613
// начало решения

// Map - карта "ключ-значение".
type Map[K comparable, V any] struct {
	Mappa map[K]V
}

// Set устанавливает значение для ключа.
func (m *Map[K, V]) Set(key K, val V) {
	if m.Mappa == nil {
		m.Mappa = make(map[K]V)
	}

	m.Mappa[key] = val
}

// Get возвращает значение по ключу.
func (m *Map[K, V]) Get(key K) V {
	return m.Mappa[key]
}

// Keys возвращает срез ключей карты.
// Порядок ключей неважен, и не обязан совпадать
// с порядком значений из метода Values.
func (m *Map[K, V]) Keys() []K {
	var res []K
	for key, _ := range m.Mappa {
		res = append(res, key)
	}
	return res
}

// Values возвращает срез значений карты.
// Порядок значений неважен, и не обязан совпадать
// с порядком ключей из метода Keys.
func (m *Map[K, V]) Values() []V {
	var res []V
	for _, val := range m.Mappa {
		res = append(res, val)
	}
	return res
}

// конец решения

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
