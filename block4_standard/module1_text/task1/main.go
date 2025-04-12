package main

import (
	"fmt"
	"strings"
)

func main() {
	src := "debugging-go-code--a-status-report"

	allow := "abcdefghijklmnopqrstuvwxyz1234567890-"

	s := strings.ToLower(src)
	s = strings.ReplaceAll(s, " ", "-")

	chars := strings.Split(s, "")

	//symbols := []string{}

	newChars := []string{}

	fmt.Println(chars)
	prev := false

	for idx, char := range chars {
		if !strings.Contains(allow, char) {
			//symbols = append(symbols, string(char))
			if !prev {
				chars[idx] = "-"
				newChars = append(newChars, "-")
			}
			prev = true

		} else {

			if !(prev && char == "-") {
				newChars = append(newChars, char)
			}
			prev = false
			if char == "-" {
				prev = true
			}

		}
	}

	fmt.Println(newChars)

	s = strings.Join(newChars, "")

	s = strings.Trim(s, "-")

	fmt.Println(s)
}
