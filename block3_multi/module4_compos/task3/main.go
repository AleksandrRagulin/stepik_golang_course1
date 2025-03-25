package main

import (
	"fmt"
	"math/rand"
)

// начало решения

type myPair struct {
	From string
	To   string
}

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for {
			select {
			case out <- randomWord(5):
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func HasUniqueLetters(s string) bool {
	// Create a map to store characters we've seen
	charMap := make(map[rune]bool)

	// Iterate over each character in the string
	for _, char := range s {
		// If the character is already in the map, it's not unique
		if charMap[char] {
			return false
		}
		// Mark the character as seen
		charMap[char] = true
	}

	// If we get here, all characters are unique
	return true
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for {
			select {
			case word, ok := <-in:
				if !ok {
					return
				}
				if HasUniqueLetters(word) {
					select {
					case out <- word:
					case <-cancel:
						return
					}
				}
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func myReverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan myPair {
	out := make(chan myPair)

	go func() {
		defer close(out)
		for {
			select {
			case word, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- myPair{word, myReverse(word)}:
				case <-cancel:
					return
				}
			case <-cancel:
				return
			}
		}
	}()

	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan myPair) <-chan myPair {
	out := make(chan myPair)

	go func() {
		defer close(out)
		for c1 != nil || c2 != nil {
			select {
			case pair, ok := <-c1:
				if !ok {
					c1 = nil
					continue
				}
				select {
				case out <- pair:
				case <-cancel:
					return
				}
			case pair, ok := <-c2:
				if !ok {
					c2 = nil
					continue
				}
				select {
				case out <- pair:
				case <-cancel:
					return
				}
			case <-cancel:
				return
			}
		}
	}()

	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan myPair, n int) {
	for i := 0; i < n; i++ {
		select {
		case out, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(out.From, "->", out.To)
		case <-cancel:
			return
		}
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
