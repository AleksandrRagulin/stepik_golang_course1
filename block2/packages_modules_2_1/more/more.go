package main

import (
	"fmt"

	"github.com/gothanks/more/numbers"
)

func main() {
	digits := numbers.AsDigits(42513)
	fmt.Println("42513 →", digits)
}
