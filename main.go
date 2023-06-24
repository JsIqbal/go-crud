package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"ISBN"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(movies)
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.Id = strconv.Itoa((rand.Intn(1000000000)))

	movies = append(movies, movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			_ = json.NewDecoder(req.Body).Decode(&movie)

			movie.Id = strconv.Itoa((rand.Intn(1000000000)))

			movies = append(movies, movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "11111", Title: "movie one", Director: &Director{FirstName: "iqbal", LastName: "hossain"}})
	movies = append(movies, Movie{Id: "2", Isbn: "11112", Title: "movie tow", Director: &Director{FirstName: "iqbal", LastName: "hossain"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8080....\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
