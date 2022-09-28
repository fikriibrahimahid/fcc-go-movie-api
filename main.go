package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Imdb     string    `json:"imdb"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

}

func createMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Imdb: "tt1745960", Title: "Top Gun: Maverick", Director: &Director{Firstname: "Joseph", Lastname: "Kosinski"}})
	movies = append(movies, Movie{ID: "2", Imdb: "tt0903747", Title: "Breaking Bad", Director: &Director{Firstname: "Vince", Lastname: "Gilligan"}})
	movies = append(movies, Movie{ID: "3", Imdb: "tt12593682", Title: "Bullet Train", Director: &Director{Firstname: "David", Lastname: "Leitch"}})
	movies = append(movies, Movie{ID: "4", Imdb: "tt9419884", Title: "Doctor Strange in the Multiverse of Madness", Director: &Director{Firstname: "Sam", Lastname: "Raimi"}})
	movies = append(movies, Movie{ID: "5", Imdb: "tt10872600", Title: "Spider-Man: No Way Home", Director: &Director{Firstname: "Jon", Lastname: "Watts"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
