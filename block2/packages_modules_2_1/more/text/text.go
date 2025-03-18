package text

import (
	"strconv"
	"strings"

	"github.com/gothanks/more/numbers"
)

// AsDigitString возвращает число в виде строки, состоящей из его цифр,
// разделенных дефисами. Например, 42513 → 4-2-5-1-3
func AsDigitString(n int) string {
	digits := numbers.AsDigits(n)
	parts := make([]string, len(digits))
	for idx, d := range digits {
		parts[idx] = strconv.Itoa(d)
	}
	return strings.Join(parts, "-")
}

// AsRunes возвращает возвращает срез с цифрами числа,
// представленными в виде рун.
func AsRunes(n int) []rune {
	return []rune(strconv.Itoa(n))
}
