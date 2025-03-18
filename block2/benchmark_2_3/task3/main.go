package main

// не удаляйте импорты, они используются при проверке
import (
	"fmt"
	"strings"
)

// Words работает со словами в строке.
type Words struct {
	subWord map[string]int
}

// MakeWords создает новый экземпляр Words.
func MakeWords(s string) Words {
	words := strings.Fields(s)
	subWord := make(map[string]int)
	for idx, word := range words {
		_, ok := subWord[word]
		if !ok {
			subWord[word] = idx
		}
	}

	return Words{subWord}
}

// Index возвращает индекс первого вхождения слова в строке,
// или -1, если слово не найдено.
func (w Words) Index(word string) int {
	idx, ok := w.subWord[word]
	if ok {
		return idx
	}
	return -1
}

func main() {
	fmt.Println("Hello World")
}
