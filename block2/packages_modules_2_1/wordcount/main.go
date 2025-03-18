package main

import (
	/*"errors"
	"flag"*/
	"fmt"
	"os"
)

func main() {
	text, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(wordCount(text))

}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (text string, err error) {
	text = os.Args[1]
	return text, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
