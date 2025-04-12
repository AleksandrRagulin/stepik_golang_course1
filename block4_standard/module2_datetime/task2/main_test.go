package main

import (
	"errors"
	"testing"
	"time"
)

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	myTime time.Time
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.myTime.Hour()
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.myTime.Minute()
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.myTime.Second()
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return t.myTime.Format("15:04:05") + " " + t.myTime.Location().String()
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	if t.myTime.Location().String() != other.myTime.Location().String() {
		return false
	}
	return t.myTime.Equal(other.myTime)
}

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.myTime.Location().String() != other.myTime.Location().String() {
		return false, errors.New("not implemented")
	}
	return t.myTime.Before(other.myTime), nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.myTime.Location().String() != other.myTime.Location().String() {
		return false, errors.New("not implemented")
	}
	return t.myTime.After(other.myTime), nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	t1 := time.Date(2017, time.January, 1, hour, min, sec, 0, loc)
	return TimeOfDay{myTime: t1}
}

// конец решения

func Test(t *testing.T) {
	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	if t1.Equal(t2) {
		t.Errorf("%v should not be equal to %v", t1, t2)
	}

	before, _ := t1.Before(t2)
	if !before {
		t.Errorf("%v should be before %v", t1, t2)
	}

	after, _ := t1.After(t2)
	if after {
		t.Errorf("%v should NOT be after %v", t1, t2)
	}
}
