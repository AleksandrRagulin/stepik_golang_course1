package main

import (
	"fmt"
	mathrand "math/rand"
	"os"
	"path/filepath"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	Total int
	files map[string]*os.File
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	return c.Total
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	c.Total++

	f, ok := c.files[string(name[0])]
	if !ok {
		filename := filepath.Join("census", string(name[0])+".txt")
		f, _ = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		c.files[string(name[0])] = f
	}

	f.WriteString(name + "\n")

}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
	for _, f := range c.files {
		f.Close()
	}
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	os.Mkdir("census", 0755)

	touch := func(path string) {
		p := filepath.FromSlash(path)
		data := []byte{}
		os.WriteFile(p, data, 0644)
	}

	runes := []rune(alphabet)

	for _, letter := range runes {
		touch(filepath.Join("census", string(letter)+".txt"))
	}

	return &Census{Total: 0, files: make(map[string]*os.File)}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

var rand = mathrand.New(mathrand.NewSource(0))

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
