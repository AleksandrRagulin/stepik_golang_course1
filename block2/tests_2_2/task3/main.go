package main

import (
	"fmt"
	"testing"
)

// WeatherServiceInterface описывает интерфейс для сервиса погоды.
type WeatherServiceInterface interface {
	Forecast() int
}

// WeatherService предсказывает погоду.
type WeatherService struct{}

// Forecast сообщает ожидаемую дневную температуру на завтра.
func (ws *WeatherService) Forecast() int {
	// магия
	return 0 // Это значение не будет использоваться в тестах
}

// Weather выдает текстовый прогноз погоды.
type Weather struct {
	service WeatherServiceInterface
}

// Forecast сообщает текстовый прогноз погоды на завтра.
func (w Weather) Forecast() string {
	deg := w.service.Forecast()
	switch {
	case deg < 10:
		return "холодно"
	case deg >= 10 && deg < 15:
		return "прохладно"
	case deg >= 15 && deg < 20:
		return "идеально"
	case deg >= 20:
		return "жарко"
	}
	return "инопланетно"
}

// MockWeatherService - заглушка для тестирования.
type MockWeatherService struct {
	deg int
}

// Forecast возвращает предопределенное значение температуры.
func (m *MockWeatherService) Forecast() int {
	return m.deg
}

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

func TestForecast(t *testing.T) {
	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			service := &MockWeatherService{deg: test.deg}
			weather := Weather{service}
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
