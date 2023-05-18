package main

import "fmt"

type Movie struct {
	Title    string
	Director string
	Year     uint
	Rating   []int
}

func AddRating(movie Movie, x int) []int {
	movie.Rating = append(movie.Rating, x)
	return movie.Rating
}
func RemoveRating(movie Movie, x int) []int {
	newrating := []int{}
	for _, rating := range movie.Rating {
		if rating != x {
			newrating = append(newrating, rating)
		}
	}
	movie.Rating = newrating
	return movie.Rating
}
func RatingAverage(movie Movie) int {
	soma := 0
	for rating := range movie.Rating {
		soma += rating
	}
	average := soma / len(movie.Rating)
	fmt.Println("Title: ", movie.Title)
	fmt.Println("Director: ", movie.Director)
	fmt.Println("Year: ", movie.Year)
	return average
}
