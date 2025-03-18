package numbers

import (
	"github.com/gothanks/more/text"
)

// IsEven проверяет число на четность.
func IsEven(n int) bool {
	return n%2 == 0
}

// AsDigits возвращает срез с цифрами числа.
func AsDigits(n int) []int {
	runes := text.AsRunes(n)
	count := len(runes)
	zero := int('0')
	digits := make([]int, count)
	for idx, char := range runes {
		digits[idx] = int(char) - zero
	}
	return digits
}
