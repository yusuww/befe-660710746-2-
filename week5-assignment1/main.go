package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Movie struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Year     int     `json:"year"`
	Rating   float64 `json:"rating"`
	Genre    string  `json:"genre"`
}


var movies = []Movie{
	{ID: "1", Title: "My Hero Academia", Director: "Kenji Nagasaki", Year: 2016, Rating: 8.7, Genre: "Anime"},
	{ID: "2", Title: "Jujutsu Kaisen", Director: "Sunghoo Park", Year: 2020, Rating: 8.9, Genre: "Anime"},
}

func getMovies(c *gin.Context) {
	yearQuery := c.Query("year")  
	genreQuery := c.Query("genre")
	directorQuery := c.Query("director") 

	filteredMovies := movies

	if yearQuery != "" {
		temp := []Movie{}
		for _, movie := range filteredMovies {
			if fmt.Sprint(movie.Year) == yearQuery {
				temp = append(temp, movie)
			}
		}
		filteredMovies = temp
	}

	if genreQuery != "" {
		temp := []Movie{}
		for _, movie := range filteredMovies {
			if movie.Genre == genreQuery {
				temp = append(temp, movie)
			}
		}
		filteredMovies = temp
	}

	if directorQuery != "" {
		temp := []Movie{}
		for _, movie := range filteredMovies {
			if movie.Director == directorQuery {
				temp = append(temp, movie)
			}
		}
		filteredMovies = temp
	}

	c.JSON(http.StatusOK, filteredMovies)
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Anime API is healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/movies", getMovies)
	}

	fmt.Println("Anime API Server starting on port 8080...")
	r.Run(":8080")
}