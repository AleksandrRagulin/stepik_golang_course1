package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {

	myInt := int(t.UnixNano())
	myString := strconv.Itoa(myInt)

	if len(myString) > 9 {
		myString = myString[:len(myString)-9] + "." + myString[len(myString)-9:]
	}

	if !strings.Contains(myString, ".") {
		myString = myString + ".0"
	}

	myString = strings.TrimRight(myString, "0")
	myString = strings.TrimRight(myString, ".")

	if !strings.Contains(myString, ".") {
		myString = myString + ".0"
	}

	return myString
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {

	re := regexp.MustCompile(`(\d+)\.(\d+)`)
	if !re.MatchString(d) {
		return time.Time{}, errors.New("invalid time")
	}

	sep := regexp.MustCompile(`\.`)
	parts := sep.Split(d, -1)

	sec, _ := strconv.Atoi(parts[0])

	if len(parts[1]) < 9 {
		count := 9 - len(parts[1])
		for i := 0; i < count; i++ {
			parts[1] = parts[1] + "0"
		}
	}

	nsec, _ := strconv.Atoi(parts[1])

	myTime := time.Date(1970, 1, 1, 0, 0, sec, nsec, time.UTC)

	return myTime, nil
}

// конец решения

func Test_asLegacyDate(t *testing.T) {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC):     "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):             "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):             "0.0",
		time.Date(2022, 5, 24, 14, 45, 22, 951205999, time.UTC): "1653403522.951205999",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123456789": time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":         time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":            time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":    time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Fatalf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}
