package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}
	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	layout := "02.01.2006"
	return time.Parse(layout, src)
}

func parseLine(line string) (time.Duration, string, error) {
	re := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)
	if !re.MatchString(line) {
		return 0, "", errors.New("invalid date")
	}

	//title
	sep := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) \s*`)
	title := sep.Split(line, -1)[1]

	//duration

	sep2 := regexp.MustCompile(`(\d+:\d+)`)
	times := sep2.FindAllString(line, 2)

	arr1 := strings.Split(times[0], ":")
	timeStr1 := fmt.Sprint(arr1[0], "h", arr1[1], "m")
	arr2 := strings.Split(times[1], ":")
	timeStr2 := fmt.Sprint(arr2[0], "h", arr2[1], "m")

	h1, _ := strconv.Atoi(arr1[0])
	h2, _ := strconv.Atoi(arr2[0])
	m1, _ := strconv.Atoi(arr1[1])
	m2, _ := strconv.Atoi(arr2[1])

	if h1 > 23 || h2 > 23 || m1 > 59 || m2 > 59 {
		return 0, "", errors.New("invalid date")
	}

	d1, _ := time.ParseDuration(timeStr1)
	d2, _ := time.ParseDuration(timeStr2)

	if d2 <= d1 {
		return 0, "", errors.New("invalid date")
	}

	return d2 - d1, title, nil
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	tmpTasks := []Task{}

	preview := make(map[string]time.Duration)

	for _, line := range lines {
		dur, title, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		preview[title] += dur
	}

	for title, dur := range preview {
		tmpTasks = append(tmpTasks, Task{date, dur, title})
	}

	return tmpTasks, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Dur > tasks[j].Dur })
}

// конец решения
// ::footer

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
