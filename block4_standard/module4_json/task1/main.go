package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	dur := time.Duration(d)
	result := dur.String()
	result = result[:len(result)-2]

	if result[:2] == "0h" {
		result = result[2:]
	}

	if result[len(result)-3:] == "h0m" {
		result = result[:len(result)-2]
	}

	b := make([]byte, 0, len(result)+2)
	b = append(b, '"')
	b = append(b, result...)
	b = append(b, '"')

	return b, nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {

	result := ""

	for i := 0; i < int(r); i++ {
		result = result + "★"
	}

	for i := 0; i < (5 - int(r)); i++ {
		result = result + "☆"
	}
	b := make([]byte, 0, len(result)+2)
	b = append(b, '"')
	b = append(b, result...)
	b = append(b, '"')
	return b, nil

}

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	result := ""
	my_indent := ""

	for i := 0; i < indent; i++ {
		my_indent += " "
	}

	if indent > 0 {
		result += "[\n"
	} else {
		result += "["
	}

	var b []byte

	if indent == 0 {
		b, _ = json.Marshal(movies)
	} else {
		b, _ = json.MarshalIndent(movies, "", my_indent)
	}

	result = string(b)

	return result, nil
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 51*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h51m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
