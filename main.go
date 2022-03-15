package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "4381", Title: "The Dark Knight Rises", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE") 
	r.HandleFunc("/movies", deleteAllMovies).Methods("DELETE") 

	fmt.Printf("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
