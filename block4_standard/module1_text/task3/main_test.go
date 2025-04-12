package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {

	newMap := []string{}

	result := ""

	newMap = append(newMap, "{")

	tmpMap := []string{}

	for k, v := range m {
		tmp := fmt.Sprintf("    %s: %d,", k, v)
		tmp = strings.ReplaceAll(tmp, `"`, "")
		tmpMap = append(tmpMap, tmp)
	}

	sort.Strings(tmpMap)

	newMap = append(newMap, tmpMap...)
	newMap = append(newMap, "}")

	if len(m) == 1 {
		result = strings.Join(newMap, " ")
		result = strings.ReplaceAll(result, ",", "")
		result = strings.ReplaceAll(result, "    ", "")
	} else if len(m) == 0 {
		result = "{}"
	} else {
		result = strings.Join(newMap, "\n")
	}

	return result
}

// конец решения

func Test1(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	got := prettify(m)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}

func Test2(t *testing.T) {
	m := map[string]int{"answer": 42}
	const want = "{ answer: 42 }"
	got := prettify(m)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}
