package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
)

type randomReader struct {
	max   int
	count int
}

func (r *randomReader) Read(p []byte) (n int, err error) {
	if r.count >= r.max {
		return 0, io.EOF
	}

	remaining := r.max - r.count
	if len(p) > remaining {
		p = p[:remaining]
	}

	n, err = rand.Read(p)
	if err != nil {
		return 0, err
	}

	r.count += n
	if r.count >= r.max {
		return n, io.EOF
	}
	return n, nil
}

func RandomReader(max int) io.Reader {
	return &randomReader{max: max}
}

// конец решения

func main() {
	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
	// (значения могут отличаться)
}
