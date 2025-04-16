package main

import (
	"fmt"
	"os"
	"strings"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {

	data, err := os.ReadFile(name)

	if err != nil {
		return nil, err
	}

	myArr := strings.Split(string(data), "\n")
	if len(myArr) >= 1 && myArr[len(myArr)-1] == "" {
		myArr = myArr[:len(myArr)-1]
	}

	return myArr, err
}

// конец решения

func main() {
	lines, err := readLines("mydata.txt")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
