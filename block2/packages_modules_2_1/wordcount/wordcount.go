package main

import "strings"

func wordCount(s string) int {
	return len(strings.Fields(s))
}
