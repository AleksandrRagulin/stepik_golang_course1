package main

import (
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {

	allow := "abcdefghijklmnopqrstuvwxyz1234567890-"

	s := strings.ToLower(src)
	s = strings.ReplaceAll(s, " ", "-")

	chars := strings.Split(s, "")

	newChars := []string{}

	prevLetter := false

	for _, char := range chars {
		if strings.Contains(allow, char) {
			if char != "-" {
				newChars = append(newChars, char)
				prevLetter = true
			} else {
				/*if prevLetter {
					newChars = append(newChars, char)
				}*/
				newChars = append(newChars, char)
				prevLetter = false
			}

		} else {
			if prevLetter {
				newChars = append(newChars, "-")
			}
			prevLetter = false
		}

	}

	s = strings.Join(newChars, "")

	s = strings.Trim(s, "-")

	return s
}

// конец решения

func Test(t *testing.T) {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func Test1(t *testing.T) {
	const phrase = "Go - Is - Awesome"
	const want = "go---is---awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func Test2(t *testing.T) {
	const phrase = "Go at Google I/O"
	const want = "go-at-google-i-o"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}

func Test3(t *testing.T) {
	const phrase = "Arrays, slices (and strings): The mechanics of 'append'"
	const want = "arrays-slices-and-strings-the-mechanics-of-append"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}
