package main

import (
	"fmt"
	"io"
	"strings"
)

// TokenReader начитывает токены из источника
type TokenReader interface {
	// ReadToken считывает очередной токен
	// Если токенов больше нет, возвращает ошибку io.EOF
	ReadToken() (string, error)
}

// TokenWriter записывает токены в приемник
type TokenWriter interface {
	// WriteToken записывает очередной токен
	WriteToken(s string) error
}

// начало решения

type MyReader struct {
	myTokens []string
}

func (reader *MyReader) ReadToken() (string, error) {
	if len(reader.myTokens) == 0 {
		return "", io.EOF
	}
	result := reader.myTokens[0]

	reader.myTokens = reader.myTokens[1:]

	return result, nil

}

type MyWriter struct {
	myTokens []string
}

func (writer *MyWriter) WriteToken(s string) error {
	writer.myTokens = append(writer.myTokens, s)
	return nil
}

func (writer *MyWriter) Words() any {
	return writer.myTokens
}

func NewWordReader(s string) *MyReader {
	return &MyReader{myTokens: strings.Split(s, " ")}
}

func NewWordWriter() *MyWriter {
	return &MyWriter{}
}

// FilterTokens читает все токены из src и записывает в dst тех,
// кто проходит проверку predicate
func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (int, error) {

	count := 0

	for {
		token, err := src.ReadToken()
		if err != nil {
			if err == io.EOF {
				break
			}
			return count, err
		}

		if predicate(token) {
			err = dst.WriteToken(token)
			if err != nil {
				return count, err
			}
			count++
		}
	}

	return count, nil
}

// конец решения

func main() {
	// Для проверки придется создать конкретные типы,
	// которые реализуют интерфейсы TokenReader и TokenWriter.

	// Ниже для примера используются NewWordReader и NewWordWriter,
	// но вы можете сделать любые на свое усмотрение.

	r := NewWordReader("go is awesome")
	w := NewWordWriter()
	predicate := func(s string) bool {
		return s != "is"
	}
	n, err := FilterTokens(w, r, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d tokens: %v\n", n, w.Words())
	// 2 tokens: [go awesome]
}
