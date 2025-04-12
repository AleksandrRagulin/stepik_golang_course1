package main

import (
	"strconv"
	"strings"
	"testing"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	result := 0

	unit := 1
	value := 0

	for _, direction := range directions {
		dir := strings.Split(direction, " ")

		if !strings.ContainsAny(direction, "1234567890") {
			continue
		}

		for _, x := range dir {
			if !strings.ContainsAny(x, "123456789") {
				continue
			}

			if strings.Contains(x, "km") {
				unit = 1000
				myData := strings.Split(x, "km")
				valuefloat, _ := strconv.ParseFloat(myData[0], 32)

				result += int(float64(unit) * valuefloat)
				break
			}

			if strings.Contains(x, "m") {
				unit = 1
				myData := strings.Split(x, "m")
				value, _ = strconv.Atoi(myData[0])

				result += unit * value
				break
			}
		}
	}

	return result
}

// конец решения

func Test(t *testing.T) {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
		"straight 1.6km",
	}
	const want = 7600
	got := calcDistance(directions)
	if got != want {
		t.Errorf("%v: got %v, want %v", directions, got, want)
	}
}
