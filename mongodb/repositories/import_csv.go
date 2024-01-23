package mongodb

import (
	"strconv"
	"strings"
)

func ImportMovieCSV(data [][]string) []Movie {
	var movies []Movie
	for i, line := range data {
		if i > 0 { // omit header line
			var movie Movie
			for j, field := range line {
				if j == 1 {
					movie.ID = field
				} else if j == 2 {
					movie.Title = field
				} else if j == 3 {
					field = strings.Trim(field, "[]")
					result := strings.Split(field, ", ")
					for i, s := range result {
						result[i] = strings.Trim(s, "'")
					}
					movie.Characters = result
				} else if j == 7 {
					field = strings.Trim(field, "[]")
					result := strings.Split(field, ", ")
					for i, s := range result {
						result[i] = strings.Trim(s, "'")
					}
					movie.Genres = result
				} else if j == 11 {
					rating, _ := strconv.ParseFloat(field, 32)
					movie.Rating = rating
				}
			}
			movies = append(movies, movie)
		}
	}
	return movies
}
