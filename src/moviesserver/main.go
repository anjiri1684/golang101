package main

import (
	"encoding/json"
	"fmt"
	"log"
	//	"math/rand"
	"net/http"
	//	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "48227",
		Title: "Movie1",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "48227",
		Title: "Movie2",
		Director: &Director{
			Firstname: "Jane",
			Lastname:  "Doe",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "48227",
		Title: "Movie3",
		Director: &Director{
			Firstname: "Vincent ",
			Lastname:  "Anjiri",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.listenAndServe(":8080", r))
}
