package main

import (
	"testing"
	"time"
)

// начало решения

func isLeapYear(year int) bool {
	t := time.Date(year, 2, 29, 17, 45, 22, 951205000, time.Local)
	return t.Month() == 2
}

// конец решения

func Test(t *testing.T) {
	if !isLeapYear(2020) {
		t.Errorf("2020 is a leap year")
	}
	if isLeapYear(2022) {
		t.Errorf("2022 is NOT a leap year")
	}
}
