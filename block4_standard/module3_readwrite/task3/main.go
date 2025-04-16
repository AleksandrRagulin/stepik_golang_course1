package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	myString, _ := reader.ReadString('\n')

	myString = myString[:len(myString)]
	myString = strings.Title(strings.ToLower(myString))

	fmt.Println(myString)
}
