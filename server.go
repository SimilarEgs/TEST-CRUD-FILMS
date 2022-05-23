package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"direcotr"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastdname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Kill Bill", Director: &Director{FirstName: "Quentin", LastName: "Tarantino"}})
	movies = append(movies, Movie{ID: "2", Title: "Fargo", Director: &Director{FirstName: "Koen", LastName: "Brothers"}})

	//Routing handlers
	//1. Get all movies
	r.HandleFunc("/movies", getMovies).Methods("GET")
	//2. Get movie by id
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	//3. Create movie
	r.HandleFunc("/movies", createMovie).Methods("POST")
	//4. Update movie by id
	r.HandleFunc("/movies/{id}", updateMovieByID).Methods("PUT")
	//5. Delete movie by id
	r.HandleFunc("/movies/{id}", deleteMoviesById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMoviesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	t := time.Now().UnixNano()
	rand.Seed(t)

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(rand.Int())
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	t := time.Now().UnixNano()
	rand.Seed(t)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			err := json.NewDecoder(r.Body).Decode(&movie)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
